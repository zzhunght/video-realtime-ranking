package mq

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(brokers []string, topic string) *KafkaProducer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:     brokers,
		Topic:       topic,
		MaxAttempts: 10,
	})

	return &KafkaProducer{
		writer: writer,
	}
}

func (p *KafkaProducer) Close() {
	if p.writer != nil {
		p.writer.Close()
	}
}

func (p *KafkaProducer) SendMessage(ctx context.Context, msg []byte) error {
	return p.writer.WriteMessages(ctx, kafka.Message{
		Value: msg,
	})
}
