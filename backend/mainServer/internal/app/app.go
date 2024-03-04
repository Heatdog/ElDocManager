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

	logger.Info("try to connect to PostgreSQL")
	postgreSQLClient, err := postgresql.NewClient(context.Background(), cfg.PostgreStorage, 5)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	defer postgreSQLClient.Close()

	logger.Info("connection to auth server")
	authConn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("auth server conn error: %s", err.Error())
	}

	authClient := authServer.NewAuthServerClient(authConn)

	res, err := authClient.GetTokens(context.Background(), &authServer.TokenRequest{
		RefreshToken: "3",
	})
	if err != nil {
		logger.Fatalf("some error %s", err.Error())
	}
	logger.Infof("Access token %s", res.AccessToken)
	logger.Infof("Refresh token %s", res.NewRefreshToken)

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
