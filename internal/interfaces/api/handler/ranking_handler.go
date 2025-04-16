package handler

import (
	"encoding/json"
	"log"
	"net/http"

	services "github.com/zzhunght/realtime-video-ranking/internal/application"
	mq "github.com/zzhunght/realtime-video-ranking/internal/infrastructure/messaging"
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

// GetVideoByRank godoc
// @Summary Get videos by rank
// @Description Fetch videos ranked by score
// @Tags ranking
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit number of videos (default: 10)"
// @Param order query string false "Sort order: ASC or DESC (default: DESC)"
// @Success 200 {object} response.Response{Data=object}
// @Failure 500 {object} response.Response{Message=string}
// @Router /api/v1/ranking/videos [get]
func (h *RankingHanlder) GetVideoByRank(w http.ResponseWriter, r *http.Request) {

	limit := readInt(r.URL.Query(), "limit", 10)
	orderby := readString(r.URL.Query(), "order", "DESC")

	reverse := true

	if orderby == "ASC" {
		reverse = false
	}
	data, err := h.rankingService.GetVideoByRank(r.Context(), limit, reverse)

	if err != nil {
		log.Printf("error when get rank video: %v \n", err)
		response.ErrorResponse(w, errors.ErrInternalServer, http.StatusInternalServerError)
		return
	}

	response.SuccessResponse(w, response.Response{
		Data: data,
	}, http.StatusOK, nil)

}

// CreateEvent godoc
// @Summary Create a new ranking event
// @Description Create a new ranking event by submitting a score for a video
// @Tags ranking
// @Accept  json
// @Produce  json
// @Param id path string true "Video ID"
// @Param payload body dto.AddScore true "Payload with score details"
// @Success 200 {object} response.Response{Data=map[string]string}
// @Failure 400 {object} response.Response{Message=string}
// @Failure 404 {object} response.Response{Message=string}
// @Failure 500 {object} response.Response{Message=string}
// @Router /api/v1/ranking/event/{id} [post]
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

	if err := payload.Validate(); err != nil {
		log.Printf("Validate fail: %v \n", err)
		response.ErrorResponse(w, errors.ErrBadrequest, http.StatusBadRequest)
		return
	}

	msg, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Err when encode data: %v \n", err)
		response.ErrorResponse(w, errors.ErrInternalServer, http.StatusInternalServerError)
		return
	}

	err = h.producer.SendMessage(r.Context(), msg)

	if err != nil {
		log.Printf("Err when send msg to kafka: %v \n", err)
		response.ErrorResponse(w, errors.ErrInternalServer, http.StatusInternalServerError)
		return
	}

	response.SuccessResponse(w, response.Response{
		Data: map[string]string{
			"message": "create event success",
		},
	}, http.StatusOK, nil)
	return

}
