package goweb

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"goutils/str"
	"net/http"
)

func (req Request) PathVariable(key string) string {
	return mux.Vars(req.Request)[key]
}

func (req Request) QueryParam(key string) string {
	return req.URL.Query().Get(key)
}

func (req Request) QueryParams(key string) []string {
	return req.URL.Query()[key]
}

func (req Request) RequestBodyFromJson(object any) error {
	return json.NewDecoder(req.Body).Decode(object)
}

func (req Request) BearerToken() string {
	authorizationValue := req.Header.Get("Authorization")
	if str_utils.IsEmpty(authorizationValue) || len(authorizationValue) < 7 {
		log.Debug().Msgf("Invalid Authorization in Header")
		return ""
	}
	bearerToken := authorizationValue[7:]
	if str_utils.IsEmpty(bearerToken) {
		log.Debug().Msgf("Empty Bearer Token in Authorization Header")
		return ""
	}
	return bearerToken
}

type Request struct {
	*http.Request
}
