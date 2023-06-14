package goweb

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

func ListenAndServe(addr string, handler http.Handler) {
	log.Info().Msgf("Listening and serving HTTP server on %s", addr)
	server := http.Server{Addr: addr, Handler: handler}
	log.Fatal().Err(server.ListenAndServe()).Msg("Server ended")
}

