package usecases

import (
	"github.com/FlyKarlik/message-service/internal/domain"
	"github.com/FlyKarlik/message-service/internal/repository"
	"github.com/FlyKarlik/message-service/internal/usecases/service"
)

type MessageUsecase interface {
	AddMessage(content string) (*domain.Message, error)
	UpdateMessageStatus(id string) error
	GetMessage(id string) (*domain.Message, error)
	GetAllMessage() ([]domain.Message, error)
	GetAllProcessedMessage() ([]domain.Message, error)
}

type StatsUsecase interface {
	GetStats() (*domain.Stats, error)
}

type Usecases struct {
	StatsUsecase   StatsUsecase
	MessageService MessageUsecase
}

func New(repo *repository.Repository) *Usecases {
	return &Usecases{
		StatsUsecase:   service.NewStatsService(repo.StatsRepository),
		MessageService: service.NewMessageService(repo.MessageRepository),
	}
}
