package goweb

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (router Router) HandlePostRequest(path string, doFunc func(Response, Request)) *mux.Route {
	return router.HandleFunc(path, convertHandlerFunction(doFunc)).Methods("POST")
}

func (router Router) HandleGetRequest(path string, doFunc func(Response, Request)) *mux.Route {
	return router.HandleFunc(path, convertHandlerFunction(doFunc)).Methods("GET")
}

func (router Router) HandlePutRequest(path string, doFunc func(Response, Request)) *mux.Route {
	return router.HandleFunc(path, convertHandlerFunction(doFunc)).Methods("PUT")
}

func (router Router) HandleDeleteRequest(path string, doFunc func(Response, Request)) *mux.Route {
	return router.HandleFunc(path, convertHandlerFunction(doFunc)).Methods("DELETE")
}

func (router Router) HandleRequest(path string, doFunc func(Response, Request)) *mux.Route {
	return router.HandleFunc(path, convertHandlerFunction(doFunc))
}

func convertHandlerFunction(f func(Response, Request)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		log.Info().Str("remote_addr", request.RemoteAddr).
			Str("method", request.Method).
			Str("url_path", request.URL.Path).
			Msg("Received a request")
		f(Response{
			ResponseWriter: response,
		}, Request{
			Request: request,
		})
	}
}

type Router struct {
	*mux.Router
}
