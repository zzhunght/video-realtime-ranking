package mq

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type ScoreConsumer struct {
	reader *kafka.Reader
}

func NewScoreConsumer(brokers []string, topic string, group string) *ScoreConsumer {

	consumer := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: group,
		// tắt tự động commit
		CommitInterval: 0,
		StartOffset:    kafka.LastOffset,
	})

	return &ScoreConsumer{
		reader: consumer,
	}
}

func (c *ScoreConsumer) Start(ctx context.Context) {

	for {
		msg, err := c.reader.ReadMessage(ctx)

		if err != nil {
			log.Printf("error when read msg : %v \n", err)
		}

		log.Printf("💬 Got message: key=%s value=%s offset=%d", string(msg.Key), string(msg.Value), msg.Offset)

		if err := c.reader.CommitMessages(ctx, msg); err != nil {
			log.Printf("Failed to commit: %v", err)
		} else {
			log.Printf("Committed offset: %d", msg.Offset)
		}
	}
}
