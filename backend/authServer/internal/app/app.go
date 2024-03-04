package app

import (
	"log"
	"net"

	"github.com/Heatdog/ElDocManager/backend/authServer/internal/server"
	authServer "github.com/Heatdog/ElDocManager/backend/authServer/pkg/proto"

	"google.golang.org/grpc"
)

func App() {
	s := grpc.NewServer()
	srv := &server.GRPCServer{}

	authServer.RegisterAuthServerServer(s, srv)

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	if err = s.Serve(l); err != nil {
		log.Fatal(err)
	}

}
