package jwttoken

import (
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenFields struct {
	ID   string
	Role string
}

func GenerateToken(fields TokenFields, key string) (string, error) {
	payload := jwt.MapClaims{
		"sub":  fields.ID,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
		"role": fields.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString(key)
}

func GenerateRefreshToken() (string, error) {
	data := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(data); err != nil {
		return "", err
	}

	return string(data), nil
}
