package goweb

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (req Request) PathVariable(key string) string {
	return mux.Vars(req.Request)[key]
}

func (req Request) QueryParam(key string) string {
	return req.URL.Query().Get(key)
}

func (req Request) RequestBodyFromJson(object any) any {
	return json.NewDecoder(req.Body).Decode(object)
}

type Request struct {
	*http.Request
}
