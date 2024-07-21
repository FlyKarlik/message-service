package service

import (
	"github.com/FlyKarlik/message-service/internal/domain"
	"github.com/FlyKarlik/message-service/internal/repository"
)

type StatsService struct {
	repo repository.StatsRepository
}

func NewStatsService(repo repository.StatsRepository) *StatsService {
	return &StatsService{repo: repo}
}

func (s *StatsService) GetStats() (*domain.Stats, error) {
	return s.repo.GetStats()
}
