package app

import (
	"ElDocManager/internal/transport"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	router := mux.NewRouter()
	router.HandleFunc("/", transport.RootHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
