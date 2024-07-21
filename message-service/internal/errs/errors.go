package errs

import "errors"

var (
	ErrServiceNameNotConfigured       = errors.New("service name field is not configured")
	ErrServerHostNotConfigured        = errors.New("server port field is not configured")
	ErrDatabaseUrlNotConfigured       = errors.New("database url field is not configured")
	ErrJaegerHostNotConfigured        = errors.New("jaeger host field is not configured")
	ErrLogLevelNotConfigured          = errors.New("log level field is not configured")
	ErrKafkaBrokersNotConfigured      = errors.New("kafka brokers field if not configured")
	ErrKafkaRequestTopicNotConfigured = errors.New("kafka request topic field is not configured")
	ErrKafkaResponseTopicNotConfigure = errors.New("kafka response topic field is not configured")
	ErrInvalidMsgId                   = errors.New("invalid message id in url param")
)
