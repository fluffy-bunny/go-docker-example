package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/gorilla/mux"
	"github.com/reugn/async"
)

func main() {
	future1 := asyncAction(8081)
	future2 := asyncAction(8082)
	_, err := future1.Join()
	if err != nil {
		log.Error().Err(err).Msg("Error while joining future1")
	}
	_, err = future2.Join()
	if err != nil {
		log.Error().Err(err).Msg("Error while joining future2")
	}
}
func asyncAction(port int) async.Future[string] {
	promise := async.NewPromise[string]()
	go func() {
		sPort := fmt.Sprintf("%d", port)
		router := mux.NewRouter()
		router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
			response := map[string]string{
				"message": "Welcome to Dockerized app",
				"port":    sPort,
			}
			json.NewEncoder(rw).Encode(response)
		})

		router.HandleFunc("/{name}", func(rw http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			name := vars["name"]
			var message string
			if name == "" {
				message = "Hello World"
			} else {
				message = "Hello " + name
			}
			response := map[string]string{
				"message": message,
				"port":    sPort,
			}
			json.NewEncoder(rw).Encode(response)
		})
		log.Info().Msgf("Server is running on port %s", sPort)
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
		if err != nil {
			log.Error().Err(err).Msg("Error while starting server")
		}
		promise.Success("OK")
	}()
	return promise.Future()

}
