package app

import (
	"fmt"
	"log"
	"net"

	"github.com/Heatdog/ElDocManager/backend/authServer/internal/config"
	"github.com/Heatdog/ElDocManager/backend/authServer/internal/server"
	"github.com/Heatdog/ElDocManager/backend/authServer/internal/server/db"
	authServer "github.com/Heatdog/ElDocManager/backend/authServer/pkg/proto"

	logger "github.com/Heatdog/ElDocManager/backend/logger/app"

	"google.golang.org/grpc"
)

func App() {

	logger := logger.GetLogger()
	cfg := config.GetConfig(logger)

	storage := db.NewRedisStorage(cfg.RedisStorage, logger, cfg.TokenKey)

	s := grpc.NewServer()
	srv := server.NewGRPCServer(storage)

	logger.Info("Register auth server")
	authServer.RegisterAuthServerServer(s, srv)

	logger.Info("listen on: ")
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.AuthStorage.BindIp, cfg.AuthStorage.Port))
	if err != nil {
		log.Fatal(err)
	}

	if err = s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
