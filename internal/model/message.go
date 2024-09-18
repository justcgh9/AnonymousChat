package model

import "time"

type Message struct {
	Id        int
	CreatedAt time.Time `db:"created_at"`
	Content   string
}