package app

import (
	"ElDocManager/internal/config"
	"ElDocManager/internal/transport"
	"ElDocManager/internal/user"
	"ElDocManager/pkg/logging"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	logger := logging.GetLogger()
	cfg := config.GetConfig()
	cors := cfg.CorsSettings()

	logger.Info("create router")
	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	logger.Info("regiser services")
	authService := user.NewAuthService()

	logger.Info("register handlers")
	handlerAuth := transport.NewHandlerAuth(logger, authService)
	handlerAuth.Register(router)

	corsHandler := cors.Handler(router)

	logger.Info("listen tcp")
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.BackendStorage.BindIp, cfg.BackendStorage.Port), corsHandler)
	if err != nil {
		log.Fatal(err)
	}

}
