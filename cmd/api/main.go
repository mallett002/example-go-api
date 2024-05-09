package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/mallett002/example-go-api/internal/handlers"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetReportCaller(true) // adds file and line# to log
	var router *chi.Mux = chi.NewRouter()
	handlers.Handler(router)

	fmt.Println("Starting GO API service...")
	err := http.ListenAndServe("localhost:8000", router)

	if err != nil {
		log.Error(err)
	}
}
