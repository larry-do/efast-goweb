package goweb

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
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
	return router.HandleFunc(path, chainBuilder(logMiddleware, router.secureFilter).build(convertToHandlerFunc(doFunc)))
}

func convertToHandlerFunc(f func(Response, Request)) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		f(Response{
			ResponseWriter: resp,
		}, Request{
			Request: req,
		})
	}
}

func (router Router) secureFilter(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		log.Debug().Str("remote_addr", req.RemoteAddr).
			Str("method", req.Method).
			Str("url_path", req.URL.Path).
			Msg("Security checking")

		if router.securityRules != nil {
			for i := range router.securityRules {
				if !router.securityRules[i](Response{ResponseWriter: resp}, Request{Request: req}) {
					return
				}
			}
		}

		next.ServeHTTP(resp, req)
	})
}

type Router struct {
	securityRules []SecurityRuleChecking
	*mux.Router
}

func NewRouter() Router {
	return Router{
		securityRules: nil,
		Router:        mux.NewRouter(),
	}
}

func (router Router) AddSecurityRules(securityRuleCheckingList ...SecurityRuleChecking) Router {
	router.securityRules = append(router.securityRules, securityRuleCheckingList...)
	return router
}
