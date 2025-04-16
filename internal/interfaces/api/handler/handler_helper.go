package handler

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"

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

func readInt(query url.Values, key string, defaultValue int) int {

	value := query.Get(key)

	if value == "" {
		return defaultValue
	}

	valueInt, err := strconv.Atoi(value)

	if err != nil {
		return defaultValue
	}

	return valueInt

}

func readString(query url.Values, key string, defaultValue string) string {

	value := query.Get(key)

	if value == "" {
		return defaultValue
	}

	return value

}
