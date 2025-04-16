package mq

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	services "github.com/zzhunght/realtime-video-ranking/internal/application"
	"github.com/zzhunght/realtime-video-ranking/internal/domain/models"
)

type ScoreConsumer struct {
	reader     *kafka.Reader
	rankingSrv *services.RankingService
}

func NewScoreConsumer(brokers []string, topic string, group string, rankingSrv *services.RankingService) *ScoreConsumer {

	consumer := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: group,
		// tắt tự động commit
		CommitInterval: 0,
	})

	return &ScoreConsumer{
		reader:     consumer,
		rankingSrv: rankingSrv,
	}
}

func (c *ScoreConsumer) Start(ctx context.Context) {
	defer c.reader.Close()
	log.Printf("Start consume event mesage.......\n")
	for {
		msg, err := c.reader.ReadMessage(ctx)

		if err != nil {
			log.Printf("error when read msg : %v \n", err)
			continue
		}

		var msgData models.Event

		err = json.Unmarshal(msg.Value, &msgData)

		if err != nil {
			log.Printf("can parse msg: %v", err)
			continue
		}

		// log.Printf("Got message: key=%s value=%v offset=%d", string(msg.Key), string(msg.Value), msg.Offset)

		score := 0

		switch msgData.Type {
		case models.ViewEvent:
			score = 1
		case models.ReactEvent:
			score = 5
		case models.CommentEvent:
			score = 10
		case models.ShareEvent:
			score = 50
		}

		if err := c.rankingSrv.AddScore(ctx, msgData.VideoID, float64(score)); err != nil {
			log.Panicf("Failed to add score: %v", err)
			continue
		}

		if err := c.reader.CommitMessages(ctx, msg); err != nil {
			log.Printf("Failed to commit: %v", err)
		} else {
			log.Printf("Committed offset: %d", msg.Offset)
		}
	}
}
