package config

import (
	"net/http"

	"github.com/rs/cors"
	"github.com/spf13/viper"
)

type CorsStrorageConfig struct {
	AllowedMethods []string
	AllowedOrigins []string
	AllowedHeader  []string
	ExposedHeaders []string
}

func CorsSettings() (*cors.Cors, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	c := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet, http.MethodDelete, http.MethodPost,
		},
		AllowedOrigins: []string{
			"http://localhost:3000",
		},
		AllowedHeaders: []string{
			"Content-Type",
		},
		ExposedHeaders: []string{
			"Authorization",
			"Content-Type",
		},
	})
	return c, nil
}
