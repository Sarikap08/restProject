package main

import (
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
	http.ListenAndServe(":5051", router)
}
