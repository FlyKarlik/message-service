package brokers

import (
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	W *kafka.Writer
	R *kafka.Reader
}

func NewKafkaClient(serviceName string, kafkaBrokers []string, requestTopic string, responseTopic string) (*KafkaClient, error) {

	writer := kafka.WriterConfig{
		Topic:    requestTopic,
		Brokers:  kafkaBrokers,
		Balancer: &kafka.LeastBytes{},
	}

	reader := kafka.ReaderConfig{
		Topic:   responseTopic,
		Brokers: kafkaBrokers,
		GroupID: serviceName,
	}

	conn, err := kafka.Dial("tcp", kafkaBrokers[0])
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(10 * time.Second))

	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             requestTopic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	})
	if err != nil {
		return nil, err
	}

	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             responseTopic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	})
	if err != nil {
		return nil, err
	}

	return &KafkaClient{W: kafka.NewWriter(writer), R: kafka.NewReader(reader)}, nil

}
