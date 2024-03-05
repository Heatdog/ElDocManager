package server

import "context"

type TokenRepository interface {
	UpdateToken(ctx context.Context, userId, role, refreshToken string) (accessToken string,
		newRefreshToken string, err error)
	InsertToken(ctx context.Context, userId, role string) (accessToken string,
		newRefreshToken string, err error)
}
