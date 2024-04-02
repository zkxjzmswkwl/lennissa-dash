package main

import (
	"fmt"
	"net/http"


	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"

	"main/internal/handlers"
	"main/internal/tools"
)

func main() {
	log.SetReportCaller(true)
	tools.NewDatabase()
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println(">.<")

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
