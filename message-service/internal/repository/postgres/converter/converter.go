package converter

import "github.com/FlyKarlik/message-service/internal/domain"

func ToMessage(dbmsg domain.DbModelMessage) *domain.Message {
	return &domain.Message{
		Id:          dbmsg.Id,
		Content:     dbmsg.Content,
		Status:      dbmsg.Status,
		CreatedAt:   dbmsg.CreatedAt,
		ProcessedAt: dbmsg.ProcessedAt.Time,
	}
}

func ToSliceMessage(dbmsg []domain.DbModelMessage) []domain.Message {

	msg := make([]domain.Message, 0, len(dbmsg))

	for i := 0; i < len(dbmsg); i++ {
		model := domain.Message{
			Id:          dbmsg[i].Id,
			Content:     dbmsg[i].Content,
			Status:      dbmsg[i].Status,
			CreatedAt:   dbmsg[i].CreatedAt,
			ProcessedAt: dbmsg[i].ProcessedAt.Time,
		}
		msg = append(msg, model)
	}

	return msg

}

func ToStats(dbstats domain.DbStats) *domain.Stats {
	return &domain.Stats{
		ProcessedCount:       dbstats.ProcessedCount,
		LastProcessedMessage: dbstats.LastProcessedMessage.String,
		LastUpdate:           dbstats.LastUpdate.Time,
	}
}
