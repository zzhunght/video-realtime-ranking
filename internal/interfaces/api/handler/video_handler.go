package handler

import "net/http"

type VideoHandler struct{}

func NewVideoHandler() *VideoHandler {
	return &VideoHandler{}
}

func (h *VideoHandler) GetVideo(w http.ResponseWriter, r *http.Request) {

}
