package goweb

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (router Router) HandlePostRequest(path string, doFunc func(Response, Request)) *mux.Route {
	return router.HandleRequest(path, doFunc).Methods("POST")
}

func (router Router) HandleGetRequest(path string, doFunc func(Response, Request)) *mux.Route {
	return router.HandleRequest(path, doFunc).Methods("GET")
}

func (router Router) HandlePutRequest(path string, doFunc func(Response, Request)) *mux.Route {
	return router.HandleRequest(path, doFunc).Methods("PUT")
}

func (router Router) HandleDeleteRequest(path string, doFunc func(Response, Request)) *mux.Route {
	return router.HandleRequest(path, doFunc).Methods("DELETE")
}

func (router Router) HandleRequest(path string, doFunc func(Response, Request)) *mux.Route {
	return router.HandleFunc(path, chainBuilder(logMiddleware, doSomethingMiddleware).build(convertHandlerFunction(doFunc)))
}

func convertHandlerFunction(f func(Response, Request)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
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
