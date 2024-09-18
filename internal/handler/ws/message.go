package ws

import (
	"log/slog"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/justcgh9/AnonymousChat/internal/model"
)

type MsgService interface {
	Create(content string) (*model.Message, error)
	GetAll() ([]*model.Message, error)
}

type WsMsgHandler struct {
	MsgListener
	WsManager  *websocket.Upgrader
	MsgService MsgService
}

type MsgListener struct {
	Clients map[*websocket.Conn]chan *model.Message // Map of clients and their channels
	mu      sync.Mutex                              // Mutex to protect Clients map
}

func NewHandler(msgService MsgService) *WsMsgHandler {
	msgL := MsgListener{
		Clients: make(map[*websocket.Conn]chan *model.Message),
	}
	wsM := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	return &WsMsgHandler{
		MsgListener: msgL,
		WsManager:   &wsM,
		MsgService:  msgService,
	}
}

func (h *WsMsgHandler) HandleConnection(w http.ResponseWriter, r *http.Request) error {
	ws, err := h.WsManager.Upgrade(w, r, nil)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	defer ws.Close()

	h.mu.Lock()
	slog.Info("Locked mutex")
	h.Clients[ws] = make(chan *model.Message, 256)
	h.mu.Unlock()
	slog.Info("Unlocked mutex")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		h.ReadMessage(r, ws)
	}()
	go func() {
		defer wg.Done()
		h.WriteMessage(r, ws)
	}()

	wg.Wait()

	slog.Info("Quitting the conenction")

	return nil
}

func (h *WsMsgHandler) ReadMessage(r *http.Request, ws *websocket.Conn) {
	slog.Info("Read Message invoked")
	for {

		msgT, content, err := ws.ReadMessage()
		if err != nil {
			slog.Error("Error reading message: " + err.Error())
			return
		}

		slog.Info("Received message", slog.Attr{Value: slog.AnyValue(content)})

		if msgT != websocket.TextMessage {
			continue
		}

		contentStr := string(content)
		if contentStr == "" {
			continue
		}

		msg, err := h.MsgService.Create(contentStr)
		if err != nil {
			slog.Error("Error creating message: " + err.Error())
			continue
		}

		h.mu.Lock()
		for client, ch := range h.Clients {
			select {
			case ch <- msg:
				slog.Info("Message broadcasted to client", slog.Attr{Key: "Client", Value: slog.AnyValue(client.RemoteAddr())})
			default:
				slog.Info("Client's channel is full, skipping", slog.Attr{Key: "Client", Value: slog.AnyValue(client.RemoteAddr())})
			}
		}
		h.mu.Unlock()
	}
}

func (h *WsMsgHandler) WriteMessage(r *http.Request, ws *websocket.Conn) {
	clientCh := h.Clients[ws]
	slog.Info("Write Message invoked")

	for {
		select {
		case msg, ok := <-clientCh:
			if !ok {
				slog.Error("Client channel closed")
				return
			}

			err := ws.WriteJSON(msg)
			if err != nil {
				slog.Error("Error writing message to client", slog.Attr{Key: "Client", Value: slog.AnyValue(ws.RemoteAddr())})
				return
			}
			slog.Info("Message sent to client", slog.Attr{Key: "Message", Value: slog.AnyValue(msg.Content)})

		case <-r.Context().Done():
			slog.Info("Connection context done")
			return
		}
	}
}
