package server

import "context"

type TokenRepository interface {
	UpdateToken(ctx context.Context, userId, refreshToken string) (accessToken string,
		newRefreshToken string, err error)
	InsertToken(ctx context.Context, userId string) (accessToken string,
		newRefreshToken string, err error)
}
