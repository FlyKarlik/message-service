package config

import (
	"os"

	"github.com/FlyKarlik/message-service/internal/errs"
)

type Config struct {
	ServiceName        string
	ServerHost         string
	DatabaseURL        string
	JaegerHost         string
	KafkaBrokers       string
	KafkaRequestTopic  string
	KafkaResponseTopic string
	LogLevel           string
}

func New() (*Config, error) {

	cfg := Config{
		ServiceName:        os.Getenv("SERVICE_NAME"),
		ServerHost:         os.Getenv("SERVER_HOST"),
		DatabaseURL:        os.Getenv("DATABASE_URL"),
		JaegerHost:         os.Getenv("JAEGER_HOST"),
		LogLevel:           os.Getenv("LOG_LEVEL"),
		KafkaBrokers:       os.Getenv("KAFKA_BROKERS"),
		KafkaRequestTopic:  os.Getenv("KAFKA_REQUEST_TOPIC"),
		KafkaResponseTopic: os.Getenv("KAFKA_RESPONSE_TOPIC"),
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c Config) Validate() error {

	configMap := map[string]struct {
		value string
		err   error
	}{
		"ServiceName":        {c.ServiceName, errs.ErrServiceNameNotConfigured},
		"ServerHost":         {c.ServerHost, errs.ErrServerHostNotConfigured},
		"DatabaseUrl":        {c.DatabaseURL, errs.ErrDatabaseUrlNotConfigured},
		"JaegerHost":         {c.JaegerHost, errs.ErrJaegerHostNotConfigured},
		"LogLevel":           {c.LogLevel, errs.ErrLogLevelNotConfigured},
		"KafkaBrokers":       {c.KafkaBrokers, errs.ErrKafkaBrokersNotConfigured},
		"KafkaRequestTopic":  {c.KafkaRequestTopic, errs.ErrKafkaRequestTopicNotConfigured},
		"KafkaResponseTopic": {c.KafkaResponseTopic, errs.ErrKafkaResponseTopicNotConfigure},
	}

	for _, val := range configMap {
		if len(val.value) == 0 {
			return val.err
		}
	}

	return nil
}
