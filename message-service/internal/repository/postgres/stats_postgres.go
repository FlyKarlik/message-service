package postgres

import (
	"github.com/FlyKarlik/message-service/internal/domain"
	"github.com/FlyKarlik/message-service/internal/repository/postgres/converter"
	"github.com/FlyKarlik/message-service/internal/repository/postgres/references"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type StatsPostgres struct {
	db *sqlx.DB
}

func NewStatsPostgres(db *sqlx.DB) *StatsPostgres {
	return &StatsPostgres{db: db}
}

func (s *StatsPostgres) GetStats() (*domain.Stats, error) {

	var model domain.DbStats

	if err := s.db.Get(&model, references.GetStatsQuery); err != nil {
		return nil, err
	}

	return converter.ToStats(model), nil
}
