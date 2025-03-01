package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
)

// NewWriter создает новый writer для указанного топика
func NewWriter(broker string, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// Produce отправляет сообщение в топик
func Produce(writer *kafka.Writer, message []byte) error {
	return writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: message,
		},
	)
}
