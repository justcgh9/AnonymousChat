package message

import "github.com/justcgh9/AnonymousChat/internal/model"

type MessageRepo interface {
	Create(content string) (*model.Message, error)
	GetAll() ([]*model.Message, error)
}

type MessageService struct {
	MsgRepo MessageRepo
}

func NewService(msgRepo MessageRepo) *MessageService {
	return &MessageService{MsgRepo: msgRepo}
}