package handler

import (
	"encoding/json"
	"log"
	"net/http"

	services "github.com/zzhunght/realtime-video-ranking/internal/application"
	mq "github.com/zzhunght/realtime-video-ranking/internal/infrastructure/mesaging"
	"github.com/zzhunght/realtime-video-ranking/internal/interfaces/api/dto"
	response "github.com/zzhunght/realtime-video-ranking/internal/interfaces/api/responses"
	errors "github.com/zzhunght/realtime-video-ranking/pkg"
)

type RankingHanlder struct {
	rankingService *services.RankingService
	producer       *mq.KafkaProducer
}

func NewRankingHanlder(rankingService *services.RankingService, producer *mq.KafkaProducer) *RankingHanlder {
	return &RankingHanlder{
		rankingService: rankingService,
		producer:       producer,
	}
}

func (h *RankingHanlder) GetVideoByRank(w http.ResponseWriter, r *http.Request) {

}

func (h *RankingHanlder) CreateEvent(w http.ResponseWriter, r *http.Request) {

	id, err := readIDParam(r)

	if err != nil {
		response.ErrorResponse(w, errors.ErrNotfound, http.StatusNotFound)
		return
	}

	payload := dto.AddScore{
		VideoID: id,
	}

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&payload)

	if err != nil {
		log.Printf("Can decode body: %v \n", err)
		response.ErrorResponse(w, errors.ErrBadrequest, http.StatusBadRequest)
		return
	}

	msg, err := json.Marshal(payload)
	err = h.producer.SendMessage(r.Context(), msg)

	if err != nil {
		log.Printf("Err when send msg to kafka: %v \n", err)
		response.ErrorResponse(w, errors.ErrInternalServer, http.StatusInternalServerError)
		return
	}

	response.SuccessResponse(w, response.Response{
		Data: map[string]string{
			"message": "send event to kafka success",
		},
	}, http.StatusOK, nil)
	return

}
