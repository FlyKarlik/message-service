package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/FlyKarlik/message-service/internal/config"
	asynchandler "github.com/FlyKarlik/message-service/internal/gateways/async/kafka/asynchandle"
	"github.com/FlyKarlik/message-service/internal/gateways/async/kafka/consumer"
	"github.com/FlyKarlik/message-service/internal/gateways/sync/http/server"
	"github.com/FlyKarlik/message-service/internal/gateways/sync/http/synchandle"
	"github.com/FlyKarlik/message-service/internal/repository"
	"github.com/FlyKarlik/message-service/internal/usecases"
	"github.com/FlyKarlik/message-service/pkg/brokers"
	"github.com/FlyKarlik/message-service/pkg/database"
	"github.com/FlyKarlik/message-service/pkg/logger"
	"github.com/FlyKarlik/message-service/pkg/tracer"
	_ "github.com/lib/pq"
)

type App struct {
	cfg *config.Config
}

func New(cfg *config.Config) *App {
	return &App{cfg: cfg}
}

// @title Message-service
// @version 1.0
// @description API server for Message-service

// @host 87.228.17.183:3000
// @BasePath /
func (a *App) Run() error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log := logger.New(a.cfg.LogLevel)

	tracer, closer, err := tracer.New(a.cfg.ServiceName, a.cfg.JaegerHost)
	if err != nil {
		return fmt.Errorf("failed to init jaeger tracer: %w", err)
	}
	defer closer.Close()

	db, err := database.ConnectToPostgres(a.cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("failed connection with postgresql: %w", err)
	}

	client, err := brokers.NewKafkaClient(a.cfg.ServiceName, []string{a.cfg.KafkaBrokers}, a.cfg.KafkaRequestTopic, a.cfg.KafkaResponseTopic)
	if err != nil {
		return fmt.Errorf("failed create kafka client: %w", err)
	}

	repo := repository.New(db)
	usecases := usecases.New(repo)
	syncHandle := synchandle.New(a.cfg, usecases, log, client.W, tracer)
	asyncHandle := asynchandler.New(a.cfg, usecases, client.R, log)

	srv := server.NewServer(a.cfg, syncHandle, log)
	consumer := consumer.NewMessageConsumerGroup([]string{a.cfg.KafkaBrokers}, "messages", client.R, log, asyncHandle)

	var wg sync.WaitGroup

	wg.Add(2)

	go srv.Run(&wg)
	go consumer.RunKafkaConsumer(&wg, ctx, cancel)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Info("Application Shutting Down")

	if err := srv.Shuttdown(ctx); err != nil {
		return fmt.Errorf("failed shuttdown http server")
	}

	if err := db.Close(); err != nil {
		return fmt.Errorf("failed close connection with postgresql: %w", err)
	}

	if err := client.W.Close(); err != nil {
		return fmt.Errorf("failed close consumer connection: %w", err)
	}

	if err := client.R.Close(); err != nil {
		return fmt.Errorf("failed close producer connection: %w", err)
	}

	return nil
}
