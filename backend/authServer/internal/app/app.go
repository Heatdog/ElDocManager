package app

import (
	"authServer/internal/server"
	authServer "authServer/pkg/proto"
	"log"
	"net"

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
