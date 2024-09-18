package message

import (
	"errors"
	"log/slog"

	"github.com/justcgh9/AnonymousChat/internal/model"
)

func (service *MessageService) Create(content string) (*model.Message, error) {
	msg, err := service.MsgRepo.Create(content)
	if err != nil {
		slog.Error(err.Error())
		return nil, errors.New("can not create message")
	}
	return msg, nil
}
