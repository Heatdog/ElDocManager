package config

import (
	"github.com/rs/cors"
)

func CorsSettings(corsConfig CorsStorageConfig) *cors.Cors {

	c := cors.New(cors.Options{
		AllowedMethods: corsConfig.AllowedMethods,
		AllowedOrigins: corsConfig.AllowedOrigins,
		AllowedHeaders: corsConfig.AllowedHeader,
		ExposedHeaders: corsConfig.ExposedHeaders,
	})
	return c
}
