package goweb

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"net/http"
	"os"
)

var config webConfig

func LoadConfiguration(configFile string) {
	log.Info().Msgf("Reading web configuration file...")
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal().Err(err).Msgf("Error loading web configuration file")
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal().Err(err).Msgf("Error parsing web configuration file")
	}
}

func listenAndServeWithAddr(addr string, handler http.Handler) {
	log.Info().Msgf("Listening and serving HTTP server on %s", addr)
	server := http.Server{Addr: addr, Handler: handler}
	log.Fatal().Err(server.ListenAndServe()).Msg("Server ended")
}

func ListenAndServe(handler http.Handler) {
	if &config == nil {
		log.Fatal().Msg("Web config not initialized.")
		return
	}
	addr := fmt.Sprintf(":%d", config.Port)
	listenAndServeWithAddr(addr, handler)
}

