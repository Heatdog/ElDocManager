package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Heatdog/ElDocManager/backend/authServer/internal/config"
	"github.com/Heatdog/ElDocManager/backend/authServer/internal/server"
	jwttoken "github.com/Heatdog/ElDocManager/backend/authServer/pkg/jwtToken"

	logger "github.com/Heatdog/ElDocManager/backend/logger/app"
	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	storage   *redis.Client
	logger    *logger.Logger
	jwtKey    string
	jwtExpire int
}

func NewRedisStorage(config config.RedisStorage, logger *logger.Logger, jwtKey string) server.TokenRepository {
	host := fmt.Sprintf("%s:%s", config.BindIp, config.Port)
	logger.Infof("redis connectin: %s", host)
	storage := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: config.Password,
		DB:       0,
	})

	if _, err := storage.Ping(context.Background()).Result(); err != nil {
		logger.Fatalf("redis connection error: %s", err.Error())
		return nil
	}

	return &RedisStorage{
		storage:   storage,
		logger:    logger,
		jwtKey:    jwtKey,
		jwtExpire: config.TokenExpiration,
	}
}

// UpdateToken - у пользователя есть рефреш токен, проверяем его. Если такого нет в базе, то ничего не отдаем.
// Если токен не совпадает, также ничего не отдаем. Если все хорошо - генерируем новый рефреш + аксес +
// сохраняем в бд

func (rs *RedisStorage) UpdateToken(ctx context.Context, userId, role, refreshToken string) (accessToken string,
	newRefreshToken string, err error) {

	storageRefreshToken, err := rs.storage.Get(ctx, userId).Result()
	if err != nil {
		rs.logger.Infof("dosn`t storage refresh token for user: %s", userId)
		rs.logger.Info(err.Error())
		return "", "", err
	}

	if strings.Compare(storageRefreshToken, refreshToken) != 0 {
		err = fmt.Errorf("refresh token dosn`t match for user: %s", userId)
		rs.logger.Info(err.Error())
		return "", "", err
	}

	acToken, err := jwttoken.GenerateToken(jwttoken.TokenFields{
		ID:   userId,
		Role: role,
	}, rs.jwtKey)
	if err != nil {
		rs.logger.Errorf("access token generation error: %s", err.Error())
		return "", "", err
	}

	rfToken, err := jwttoken.GenerateRefreshToken()
	if err != nil {
		rs.logger.Errorf("refresh token generation error: %s", err.Error())
		return "", "", nil
	}

	expire := time.Hour * 60 * time.Duration(rs.jwtExpire)

	if err = rs.storage.Set(ctx, userId, rfToken, expire).Err(); err != nil {
		rs.logger.Errorf("redis save error: %s", err.Error())
		return "", "", nil
	}

	return acToken, rfToken, nil
}

// У пользователя нет токена - добавляем в базу.

func (rs *RedisStorage) InsertToken(ctx context.Context, userId, role string) (accessToken string,
	newRefreshToken string, err error) {

	rfToken, err := jwttoken.GenerateRefreshToken()
	if err != nil {
		rs.logger.Errorf("refresh token generation error: %s", err.Error())
		return "", "", nil
	}

	expire := time.Hour * 60 * time.Duration(rs.jwtExpire)

	if err = rs.storage.Set(ctx, userId, rfToken, expire).Err(); err != nil {
		rs.logger.Errorf("redis save error: %s", err.Error())
		return "", "", err
	}

	acToken, err := jwttoken.GenerateToken(jwttoken.TokenFields{
		ID:   userId,
		Role: role,
	}, rs.jwtKey)
	if err != nil {
		rs.logger.Errorf("access token generation error: %s", err.Error())
		return "", "", err
	}

	return acToken, rfToken, nil
}
