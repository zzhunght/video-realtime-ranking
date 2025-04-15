package handler

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func readIDParam(r *http.Request) (id string, err error) {

	params := httprouter.ParamsFromContext(r.Context())

	id = params.ByName("id")

	if id == "" {
		return "", errors.New("invalid parameter")
	}
	return
}
