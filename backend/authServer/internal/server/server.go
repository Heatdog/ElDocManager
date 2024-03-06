package server

import (
	"context"

	authServer "github.com/Heatdog/ElDocManager/AuthServer/pkg/proto"
)

type GRPCServer struct {
	authServer.UnimplementedAuthServerServer
	storage TokenRepository
}

func NewGRPCServer(repo TokenRepository) *GRPCServer {
	return &GRPCServer{
		storage: repo,
	}
}

type RefreshTokenUserId struct {
	ResfreshToken string
	UserId        string
}

func (s *GRPCServer) GetTokens(ctx context.Context, req *authServer.TokenRequest) (*authServer.TokenRespons, error) {

	accessToken, refreashToken, err := s.storage.UpdateToken(ctx, req.UserId, req.Role, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &authServer.TokenRespons{
		AccessToken:     accessToken,
		NewRefreshToken: refreashToken,
	}, nil
}

func (s *GRPCServer) CreateRefreshToken(ctx context.Context,
	req *authServer.TokenCreateRequest) (*authServer.TokenRespons, error) {

	accessToken, refreshToken, err := s.storage.InsertToken(ctx, req.UserId, req.Role)
	if err != nil {
		return nil, err
	}

	return &authServer.TokenRespons{
		AccessToken:     accessToken,
		NewRefreshToken: refreshToken,
	}, nil
}
