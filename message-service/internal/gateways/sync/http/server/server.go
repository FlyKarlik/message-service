package server

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/FlyKarlik/message-service/internal/config"
	"github.com/FlyKarlik/message-service/internal/gateways/sync/http/synchandle"
	"github.com/FlyKarlik/message-service/pkg/logger"
)

type Server struct {
	cfg        *config.Config
	log        *logger.Logger
	handler    *synchandle.SyncHandler
	httpserver *http.Server
}

func NewServer(cfg *config.Config, handler *synchandle.SyncHandler, log *logger.Logger) *Server {
	return &Server{cfg: cfg, handler: handler, log: log}
}

func (s *Server) Run(wg *sync.WaitGroup) {

	defer wg.Done()

	handler := NewRouter(s.handler)

	s.httpserver = &http.Server{
		Addr:              s.cfg.ServerHost,
		Handler:           handler,
		MaxHeaderBytes:    1 << 20,
		ReadTimeout:       time.Minute,
		WriteTimeout:      time.Minute,
		ReadHeaderTimeout: time.Minute,
	}

	err := s.httpserver.ListenAndServe()
	if err != nil {
		s.log.Error("can not serve", "error", err)
		return
	}
}

func (s *Server) Shuttdown(ctx context.Context) error {
	return s.httpserver.Shutdown(ctx)
}
