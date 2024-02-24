package app

import (
	"ElDocManager/internal/auth"
	"ElDocManager/internal/config"
	"ElDocManager/pkg/logging"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	logger := logging.GetLogger()

	logger.Info("create router")
	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	logger.Info("register auth handler")
	handler := auth.NewHandler(logger)
	handler.Register(router)

	cors, err := config.CorsSettings()
	if err != nil {
		log.Fatalf(err.Error())
	}

	corsHandler := cors.Handler(router)

	logger.Info("start application")
	err = http.ListenAndServe("127.0.0.1:8080", corsHandler)
	if err != nil {
		log.Fatal(err)
	}

}
