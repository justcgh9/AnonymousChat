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

type MsgListener struct {
	MsgCh chan *model.Message
}

type WsMsgHandler struct {
	MsgListener
	WsManager   *websocket.Upgrader
	MsgService	MsgService
}

func NewHandler(msgService MsgService) *WsMsgHandler {
	msgL := MsgListener{MsgCh: make(chan *model.Message)}
	wsM := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {return true}}
	return &WsMsgHandler{MsgListener: msgL, WsManager: &wsM, MsgService: msgService}
}

func (h *WsMsgHandler) HandleConnection(w http.ResponseWriter, r *http.Request) error {
	ws, err := h.WsManager.Upgrade(w, r, nil)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	defer ws.Close()

	var wg sync.WaitGroup
	wg.Add(2)
	
	go func(){
		defer wg.Done()
		h.ReadMessage(r, ws)
	}()
	go func() {
		defer wg.Done()
		h.WriteMessage(r, ws)
	}()

	wg.Wait()

	return nil
}

func (h *WsMsgHandler) ReadMessage(r *http.Request, ws *websocket.Conn) {
	for {
		msgT, content, err := ws.ReadMessage()
		if err != nil {
			slog.Error(err.Error())
			return
		}
		if msgT != websocket.TextMessage {
			continue
		}
		
		contentStr := string(content)
		if contentStr == "" {
			continue
		}
		
		msg, err := h.MsgService.Create(string(content))
		if err != nil {
			slog.Error(err.Error())
			return
		}

		select {
			case h.MsgCh<- msg:
				slog.Info("message is sent")
			case <-r.Context().Done():
				return
		}
	}
}

func (h *WsMsgHandler) WriteMessage(r *http.Request, ws *websocket.Conn) {
	for {
		select {
			case msg, ok := <-h.MsgCh:
				if !ok {
					return
				}
				slog.Info("message is recieved")

				ws.WriteMessage(websocket.TextMessage, []byte(msg.Content))
			case <-r.Context().Done():
				return
		}
	}
}
