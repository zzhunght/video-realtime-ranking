package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	router *httprouter.Router
}

func NewRouter() *Router {
	return &Router{
		router: httprouter.New(),
	}
}
func healthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (r *Router) SetupRouter() http.Handler {
	r.router.GET("/health-check", healthCheck)
	return r.router
}
