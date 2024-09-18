package message

import "github.com/justcgh9/AnonymousChat/internal/model"

func (repo *MsgRepo) Create(content string) (*model.Message, error) {
	msg := model.Message{}
	err := repo.DB.Get(&msg, "INSERT INTO message (content) VALUES($1) RETURNING *", content)
	return &msg, err
}
