package main

import (
	"golangapi/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HandlerSendMeals)

	log.Fatal(http.ListenAndServe(":8080", r))
}
