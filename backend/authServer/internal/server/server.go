package server

import (
	"context"
	"fmt"

	authServer "github.com/Heatdog/ElDocManager/backend/authServer/pkg/proto"

	"github.com/redis/go-redis/v9"
)

type GRPCServer struct {
	authServer.UnimplementedAuthServerServer
}

type RefreshTokenUserId struct {
	ResfreshToken string
	UserId        string
}

func (s *GRPCServer) GetTokens(ctx context.Context, req *authServer.TokenRequest) (*authServer.TokenRespons, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ping, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}

	fmt.Println(ping)

	if err = client.Set(ctx, "name", "Elliot", 0).Err(); err != nil {
		fmt.Printf("Failed to set value in the redis instance: %s", err.Error())
		return nil, nil
	}

	val, err := client.Get(ctx, "name").Result()
	if err != nil {
		fmt.Printf("failed to get value from redis: %s", err.Error())
		return nil, nil
	}

	fmt.Printf("value from redis from redis: %s", val)

	return &authServer.TokenRespons{
		AccessToken:     "1",
		NewRefreshToken: "2",
	}, nil
}

func (s *GRPCServer) CreateRefreshToken(ctx context.Context,
	req *authServer.TokenCreateRequest) (*authServer.TokenRespons, error) {
	return &authServer.TokenRespons{
		AccessToken:     "1",
		NewRefreshToken: "2",
	}, nil
}
