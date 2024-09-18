package message

import (
	"errors"
	"log/slog"

	"github.com/justcgh9/AnonymousChat/internal/model"
)

func (service *MessageService) GetAll() ([]*model.Message, error) {
	msgs, err := service.MsgRepo.GetAll()
	if err != nil {
		slog.Error(err.Error())
		return nil, errors.New("can not get messages")
	}
	return msgs, nil
}
