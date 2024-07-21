package repository

import (
	"github.com/FlyKarlik/message-service/internal/domain"
	"github.com/FlyKarlik/message-service/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type MessageRepository interface {
	AddMessage(content string) (*domain.Message, error)
	UpdateMessageStatus(id string) error
	GetMessage(id string) (*domain.Message, error)
	GetAllMessage() ([]domain.Message, error)
	GetAllProcessedMessage() ([]domain.Message, error)
}

type StatsRepository interface {
	GetStats() (*domain.Stats, error)
}

type Repository struct {
	MessageRepository MessageRepository
	StatsRepository   StatsRepository
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		StatsRepository:   postgres.NewStatsPostgres(db),
		MessageRepository: postgres.NewMessagePostgres(db),
	}
}
