package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Heatdog/ElDocManager/backend/mainServer/internal/config"
	"github.com/Heatdog/ElDocManager/backend/mainServer/internal/transport"
	"github.com/Heatdog/ElDocManager/backend/mainServer/internal/user"
	userDb "github.com/Heatdog/ElDocManager/backend/mainServer/internal/user/db"
	"github.com/Heatdog/ElDocManager/backend/mainServer/pkg/client/postgresql"
	"github.com/Heatdog/ElDocManager/backend/mainServer/pkg/logging"

	authServer "github.com/Heatdog/ElDocManager/backend/authServer/pkg/proto"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func Run() {
	logger := logging.GetLogger()
	cfg := config.GetConfig(logger)
	cors := cfg.CorsSettings()
	ctx := context.Background()

	logger.Info("try to connect to PostgreSQL")
	postgreSQLClient, err := postgresql.NewClient(ctx, cfg.PostgreStorage, 5)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	defer postgreSQLClient.Close()

	logger.Info("connection to auth server")
	authConn, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.AuthServerStorage.BindIp, cfg.AuthServerStorage.Port),
		grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("auth server conn error: %s", err.Error())
	}

	authServer.NewAuthServerClient(authConn)

	logger.Info("Init repos")
	userRepo := userDb.NewUserRepository(postgreSQLClient, logger)

	logger.Info("create router")
	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	logger.Info("regiser services")
	userService := user.NewUserService(logger, userRepo)

	logger.Info("register handlers")
	handlerUser := transport.NewUserHandler(logger, userService, cfg.JwtKey)
	handlerUser.Register(router)

	corsHandler := cors.Handler(router)

	logger.Info("listen tcp")
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.BackendStorage.BindIp, cfg.BackendStorage.Port), corsHandler)
	if err != nil {
		log.Fatal(err)
	}

}
