package postgres

import (
	"time"

	"github.com/FlyKarlik/message-service/internal/domain"
	"github.com/FlyKarlik/message-service/internal/repository/postgres/converter"
	"github.com/FlyKarlik/message-service/internal/repository/postgres/references"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type MessagePostgres struct {
	db *sqlx.DB
}

func NewMessagePostgres(db *sqlx.DB) *MessagePostgres {
	return &MessagePostgres{db: db}
}

func (m *MessagePostgres) AddMessage(content string) (*domain.Message, error) {

	var model domain.DbModelMessage
	var id string

	if err := m.db.QueryRow(references.AddMessageQuery, content).Scan(&id); err != nil {
		return nil, err
	}

	if err := m.db.Get(&model, references.GetMessageQuery, id); err != nil {
		return nil, err
	}

	return converter.ToMessage(model), nil
}

func (m *MessagePostgres) GetMessage(id string) (*domain.Message, error) {

	var model domain.DbModelMessage

	if err := m.db.Get(&model, references.GetMessageQuery, id); err != nil {
		return nil, err
	}

	return converter.ToMessage(model), nil
}

func (m *MessagePostgres) GetAllMessage() ([]domain.Message, error) {

	var model []domain.DbModelMessage

	if err := m.db.Select(&model, references.GetAllMessageQuery); err != nil {
		return nil, err
	}

	return converter.ToSliceMessage(model), nil
}

func (m *MessagePostgres) GetAllProcessedMessage() ([]domain.Message, error) {

	var model []domain.DbModelMessage

	if err := m.db.Select(&model, references.GetAllProcessedMsgQuery); err != nil {
		return nil, err
	}
	return converter.ToSliceMessage(model), nil
}

func (m *MessagePostgres) UpdateMessageStatus(id string) error {
	_, err := m.db.Exec(references.UpdateStatusMsgQuery, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}
