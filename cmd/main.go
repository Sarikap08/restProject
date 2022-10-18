package main

import (
	"log"
	"net/http"

	"github.com/Sarikap08/restProject/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	buildHttpServer()
}

func buildHttpServer() {
	router := mux.NewRouter().SkipClean(true)
	handlers.NewSystemHandler().RegisterEndpoint(router)
	err := http.ListenAndServe(":5051", router)
	if err != nil {
		log.Fatal("Encounter error while starting the HTTP server : ", err)
		return
	}
}
