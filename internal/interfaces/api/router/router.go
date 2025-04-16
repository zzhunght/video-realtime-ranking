package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/zzhunght/realtime-video-ranking/docs/api"
	"github.com/zzhunght/realtime-video-ranking/internal/interfaces/api/handler"
)

type Router struct {
	router         *httprouter.Router
	rankingHandler *handler.RankingHanlder
}

func NewRouter(
	rankingHandler *handler.RankingHanlder,
) *Router {
	return &Router{
		router:         httprouter.New(),
		rankingHandler: rankingHandler,
	}
}
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (r *Router) SetupRouter() http.Handler {
	r.router.HandlerFunc(http.MethodGet, "/health-check", healthCheck)
	r.router.HandlerFunc(http.MethodGet, "/api/v1/ranking/videos", r.rankingHandler.GetVideoByRank)
	r.router.HandlerFunc(http.MethodPost, "/api/v1/ranking/event/:id", r.rankingHandler.CreateEvent)

	r.router.Handler(http.MethodGet, "/swagger/*filepath", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
	return r.router
}
