package app

import (
	"ElDocManager/internal/config"
	"ElDocManager/internal/transport"
	"ElDocManager/internal/user"
	userDb "ElDocManager/internal/user/db"
	"ElDocManager/pkg/client/postgresql"
	"ElDocManager/pkg/logging"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	logger := logging.GetLogger()
	cfg := config.GetConfig(logger)
	cors := cfg.CorsSettings()

	logger.Info("try to connect to PostgreSQL")
	postgreSQLClient, err := postgresql.NewClient(context.Background(), cfg.PostgreStorage, 5)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	logger.Info("Init repos")
	userRepo := userDb.NewUserRepository(postgreSQLClient, logger)

	userRepo.FindAll(context.Background())

	logger.Info("create router")
	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	logger.Info("regiser services")
	authService := user.NewAuthService()

	logger.Info("register handlers")
	handlerAuth := transport.NewHandlerAuth(logger, authService)
	handlerAuth.Register(router)

	corsHandler := cors.Handler(router)

	logger.Info("listen tcp")
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.BackendStorage.BindIp, cfg.BackendStorage.Port), corsHandler)
	if err != nil {
		log.Fatal(err)
	}

}
