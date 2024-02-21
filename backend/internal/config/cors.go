package config

import (
	"net/http"

	"github.com/rs/cors"
	"github.com/spf13/viper"
)

func CorsSettings() (*cors.Cors, error) {
	viper.SetConfigName("frontend")
	viper.SetConfigType("json")
	viper.AddConfigPath("./configs/frontend")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	c := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet, http.MethodDelete, http.MethodPost,
		},
		AllowedOrigins: []string{
			viper.GetString("frontendHost"),
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
