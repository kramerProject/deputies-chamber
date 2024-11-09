package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/kramerProject/deputies-chamber/application"
)

func MakeDeputiesHandlers(r *mux.Router, n *negroni.Negroni, service application.DeputiesService) {
	r.Handle("/deputies", n.With(
		negroni.Wrap(getDeputies(service)),
	)).Methods("GET", "OPTIONS")
}

func getDeputies(service application.DeputiesService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("Handleerrrr")
		deputies, err := service.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(deputies)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
