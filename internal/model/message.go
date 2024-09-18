package model

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id        uuid.UUID
	CreatedAt time.Time `db:"created_at"`
	Content   string
}