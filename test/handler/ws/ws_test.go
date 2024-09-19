package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/justcgh9/AnonymousChat/database/sqlite"
	"github.com/justcgh9/AnonymousChat/internal/handler/ws"
	"github.com/justcgh9/AnonymousChat/internal/model"
	messageRepo "github.com/justcgh9/AnonymousChat/internal/repo/message"
	messageService "github.com/justcgh9/AnonymousChat/internal/service/message"
	. "github.com/justcgh9/AnonymousChat/test/handler/common"
)

func getHandler() *ws.WsMsgHandler {
	db := sqlite.NewConn()
	msgR := messageRepo.NewRepo(db)
	msgS := messageService.NewService(msgR)
	return ws.NewHandler(msgS)
}

func getMessages() []model.Message {
	db := sqlite.NewConn()
	defer db.Close()
	messages := []model.Message{}
	db.Select(&messages, "SELECT * FROM message")

	return messages
}

func TestGetMessage(t *testing.T) {
	old := SetupDB()
	defer TearDownDb(old)

	wsH := getHandler()
	
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wsH.HandleConnection(w, r)
	}))
    defer s.Close()

    u := "ws" + strings.TrimPrefix(s.URL, "http")
    ws, _, err := websocket.DefaultDialer.Dial(u, nil)
    if err != nil {
        t.Fatalf("%v", err)
    }
    defer ws.Close()

    if err := ws.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
        t.Fatalf("%v", err)
    }
    _, p, err := ws.ReadMessage()
    if err != nil {
		t.Fatalf("%v", err)
    }
    
    msgs := getMessages()
    msg := model.Message{}
    json.Unmarshal(p, &msg)
    if msg != msgs[0] {
        t.Fatalf("bad message: %v", msg)
    }
}
