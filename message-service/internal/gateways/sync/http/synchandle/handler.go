package synchandle

import (
	"github.com/FlyKarlik/message-service/internal/config"
	"github.com/FlyKarlik/message-service/internal/usecases"
	"github.com/FlyKarlik/message-service/pkg/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"
)

type SyncHandler struct {
	cfg      *config.Config
	usecases *usecases.Usecases
	log      *logger.Logger
	w        *kafka.Writer
	trace    opentracing.Tracer
}

func New(cfg *config.Config, usecases *usecases.Usecases, log *logger.Logger, w *kafka.Writer, trace opentracing.Tracer) *SyncHandler {
	return &SyncHandler{cfg: cfg, usecases: usecases, log: log, w: w, trace: trace}
}
