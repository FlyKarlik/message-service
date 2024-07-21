package asynchandler

import (
	"github.com/FlyKarlik/message-service/internal/config"
	"github.com/FlyKarlik/message-service/internal/usecases"
	"github.com/FlyKarlik/message-service/pkg/logger"
	"github.com/segmentio/kafka-go"
)

type AsyncHandler struct {
	cfg      *config.Config
	usecases *usecases.Usecases
	r        *kafka.Reader
	log      *logger.Logger
}

func New(cfg *config.Config, usecases *usecases.Usecases, r *kafka.Reader, log *logger.Logger) *AsyncHandler {
	return &AsyncHandler{cfg: cfg, usecases: usecases, r: r, log: log}
}
