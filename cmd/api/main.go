package main

import (
	"fmt"
	"go_api/internal/handlers"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Starting go api service...")

	err := http.ListenAndServe("Localhost:7000", r)
	if err != nil {
		log.Error(err)
	}
}
