package http

import (
	"encoding/json"
	httpNet "net/http"

	"github.com/justcgh9/AnonymousChat/internal/model"
)

type MsgService interface {
	GetAll() ([]*model.Message, error)
}

type HttpMsgHandler struct {
	MsgService
}

func NewHandler(msgService MsgService) *HttpMsgHandler {
	return &HttpMsgHandler{MsgService: msgService}
}

func (h *HttpMsgHandler) HandleGetMessageCount(w httpNet.ResponseWriter, r *httpNet.Request) error {
	msgs, err := h.GetAll()
	if err != nil {
		return err
	}
	count, _ := json.Marshal(len(msgs))
	w.Write(count)
	return nil
}
