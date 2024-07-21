package asynchandler

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func (a *AsyncHandler) HandleUpdateMessageRequest(ctx context.Context, msg kafka.Message) {

	var id = msg.Headers[0].Key

	if err := a.usecases.MessageService.UpdateMessageStatus(id); err != nil {
		a.log.Error("[AsyncHandler.HandlerProcessedMessage] a.usecases.MessageService.UpdateMessageStatus failed", err)
		return
	}

	if err := a.r.CommitMessages(ctx, msg); err != nil {
		a.log.Error("[AsyncHandler] a.r.CommitMessages failed", err)
		return
	}
}
