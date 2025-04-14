package handler

import "net/http"

type RankingHanlder struct {
}

func NewRankingHanlder() *RankingHanlder {
	return &RankingHanlder{}
}

func (h *RankingHanlder) GetVideoByRank(w http.ResponseWriter, r *http.Request) {

}
