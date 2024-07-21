package consumer

import (
	"context"
	"log"
	"sync"
)

func (m *MessageConsumerGroup) messageWorker(wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc) {

	defer wg.Done()
	defer cancel()

	for {
		message, err := m.consumer.ReadMessage(ctx)
		if err != nil {
			log.Println(err)
			continue
		}

		switch message.Key[0] {
		case UpdateMessageRequest:
			m.handler.HandleUpdateMessageRequest(ctx, message)
			continue
		}
	}
}
