package message

import "github.com/justcgh9/AnonymousChat/internal/model"

func (repo *MsgRepo) GetAll() ([]*model.Message, error) {
	msgs := []*model.Message{}
	err := repo.DB.Select(&msgs, "SELECT * FROM message ORDER BY created_at")
	return msgs, err
}