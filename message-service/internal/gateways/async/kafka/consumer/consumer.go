package consumer

import (
	"context"
	"sync"

	asynchandler "github.com/FlyKarlik/message-service/internal/gateways/async/kafka/asynchandle"
	"github.com/FlyKarlik/message-service/pkg/logger"
	"github.com/segmentio/kafka-go"
)

type MessageConsumerGroup struct {
	brokers  []string
	groupId  string
	consumer *kafka.Reader
	log      *logger.Logger
	handler  *asynchandler.AsyncHandler
}

func NewMessageConsumerGroup(
	brokers []string,
	groupId string,
	consumer *kafka.Reader,
	log *logger.Logger,
	handler *asynchandler.AsyncHandler) *MessageConsumerGroup {

	return &MessageConsumerGroup{
		brokers:  brokers,
		groupId:  groupId,
		consumer: consumer,
		log:      log,
		handler:  handler,
	}
}

func (m *MessageConsumerGroup) RunKafkaConsumer(wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc) {
	go m.messageWorker(wg, ctx, cancel)
}
