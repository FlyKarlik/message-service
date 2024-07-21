package service

import (
	"github.com/FlyKarlik/message-service/internal/domain"
	"github.com/FlyKarlik/message-service/internal/repository"
)

type MessageService struct {
	repo repository.MessageRepository
}

func NewMessageService(repo repository.MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (m *MessageService) AddMessage(content string) (*domain.Message, error) {
	return m.repo.AddMessage(content)
}

func (m *MessageService) GetMessage(id string) (*domain.Message, error) {
	return m.repo.GetMessage(id)
}

func (m *MessageService) GetAllMessage() ([]domain.Message, error) {
	return m.repo.GetAllMessage()
}

func (m *MessageService) GetAllProcessedMessage() ([]domain.Message, error) {
	return m.repo.GetAllProcessedMessage()
}

func (m *MessageService) UpdateMessageStatus(id string) error {
	return m.repo.UpdateMessageStatus(id)
}
