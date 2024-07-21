package domain

import (
	"database/sql"
	"time"
)

type Stats struct {
	ProcessedCount       int       `json:"processed_count"`
	LastProcessedMessage string    `json:"last_processed_message"`
	LastUpdate           time.Time `json:"last_update"`
}

type DbStats struct {
	ProcessedCount       int            `db:"processed_count"`
	LastProcessedMessage sql.NullString `db:"last_processed_message"`
	LastUpdate           sql.NullTime   `db:"last_update"`
}
