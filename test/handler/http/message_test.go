package http

import (
	"encoding/json"
	httpNet "net/http"
	httpTest "net/http/httptest"
	"slices"
	"testing"

	"github.com/justcgh9/AnonymousChat/database/sqlite"
	httpH "github.com/justcgh9/AnonymousChat/internal/handler/http"
	"github.com/justcgh9/AnonymousChat/internal/model"
	messageRepo "github.com/justcgh9/AnonymousChat/internal/repo/message"
	messageService "github.com/justcgh9/AnonymousChat/internal/service/message"
	. "github.com/justcgh9/AnonymousChat/test/handler/common"
)

func getHandler() *httpH.HttpMsgHandler {
	db := sqlite.NewConn()
	msgR := messageRepo.NewRepo(db)
	msgS := messageService.NewService(msgR)
	return httpH.NewHandler(msgS)
}

func createMessages(n int) []model.Message {
	db := sqlite.NewConn()
	defer db.Close()

	for i := 0; i < n; i++ {
		db.Exec("INSERT INTO message(content) VALUES($1)", i)
	}

	messages := []model.Message{}
	db.Select(&messages, "SELECT * FROM message")

	return messages
}

func TestGetZeroMessageCount(t *testing.T) {
	old := SetupDB()
	defer TearDownDb(old)

	msgH := getHandler()

	req, err := httpNet.NewRequest("GET", "/messages/count", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httpTest.NewRecorder()
	handler := httpNet.HandlerFunc(func(w httpNet.ResponseWriter, r *httpNet.Request) {
		msgH.HandleGetMessagesCount(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != httpNet.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, httpNet.StatusOK)
	}

	expected := "0"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetNMessageCount(t *testing.T) {
	old := SetupDB()
	defer TearDownDb(old)

	msgH := getHandler()
	createMessages(10)

	req, err := httpNet.NewRequest("GET", "/messages/count", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httpTest.NewRecorder()
	handler := httpNet.HandlerFunc(func(w httpNet.ResponseWriter, r *httpNet.Request) {
		msgH.HandleGetMessagesCount(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != httpNet.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, httpNet.StatusOK)
	}

	expected := "10"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetMessages(t *testing.T) {
	old := SetupDB()
	defer TearDownDb(old)

	msgH := getHandler()
	msgs := createMessages(10)

	req, err := httpNet.NewRequest("GET", "/messages", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httpTest.NewRecorder()
	handler := httpNet.HandlerFunc(func(w httpNet.ResponseWriter, r *httpNet.Request) {
		msgH.HandleGetMessages(w, r)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != httpNet.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, httpNet.StatusOK)
	}

	var resmgs []model.Message
	json.NewDecoder(rr.Body).Decode(&resmgs)
	if !slices.Equal(resmgs, msgs) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), msgs)
	}
}
