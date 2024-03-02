package server

import (
	authServer "authServer/pkg/proto"
	"context"
)

type GRPCServer struct {
	authServer.UnimplementedAuthServerServer
}

func (s *GRPCServer) GetTokens(ctx context.Context, req *authServer.TokenRequest) (*authServer.TokenRespons, error) {
	return &authServer.TokenRespons{
		AccessToken:     "1",
		NewRefreshToken: "2",
	}, nil
}
