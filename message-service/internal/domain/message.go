package domain

import (
	"database/sql"
	"time"
)

type Message struct {
	Id          string    `json:"id"`
	Content     string    `json:"content" `
	Status      int       `json:"status" `
	CreatedAt   time.Time `json:"created_at" `
	ProcessedAt time.Time `json:"processed_at" `
}

type DbModelMessage struct {
	Id          string       `db:"id"`
	Content     string       `db:"content"`
	Status      int          `db:"status"`
	CreatedAt   time.Time    `db:"created_at"`
	ProcessedAt sql.NullTime `db:"processed_at"`
}
