package dto

import "time"

type MessageDto struct {
	ID        string
	Message   string
	Timestamp time.Time
}
