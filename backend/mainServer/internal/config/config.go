package config

import (
	"sync"

	"github.com/Heatdog/ElDocManager/backend/mainServer/pkg/logging"

	"github.com/rs/cors"
	"github.com/spf13/viper"
)

type Config struct {
	JwtKey            string               `yaml:"jwt_auth_key"`
	BackendStorage    ListenBackend        `mapstructure:"main_server_listen"`
	CorseStorage      CorsStorageConfig    `mapstructure:"cors_settings"`
	PostgreStorage    PostgreStorageConfig `mapstructure:"postgre_settings"`
	AuthServerStorage ListenAuthServer     `mapstructure:"auth_server_listen"`
}

type ListenBackend struct {
	Type   string `mapstructure:"type"`
	BindIp string `mapstructure:"bind_ip"`
	Port   string `mapstructure:"port"`
}

type CorsStorageConfig struct {
	AllowedMethods []string `mapstructure:"allowed_methods"`
	AllowedOrigins []string `mapstructure:"allowed_origins"`
	AllowedHeader  []string `mapstructure:"allowed_headers"`
	ExposedHeaders []string `mapstructure:"exposed_headers"`
}

type PostgreStorageConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"database"`
}

type ListenAuthServer struct {
	Type   string `mapstructure:"type"`
	BindIp string `mapstructure:"bind_ip"`
	Port   string `mapstructure:"port"`
}

var instance *Config
var once sync.Once

func GetConfig(logger *logging.Logger) *Config {
	once.Do(func() {
		logger.Info("read application instance")
		instance = &Config{}

		viper.SetConfigFile("../configs/config.yaml")
		if err := viper.ReadInConfig(); err != nil {
			logger.Fatal(err)
		}
		if err := viper.Unmarshal(instance); err != nil {
			logger.Fatal(err)
		}

		viper.SetConfigFile("../../configs/config.yaml")
		if err := viper.ReadInConfig(); err != nil {
			logger.Fatal(err)
		}
		if err := viper.Unmarshal(instance); err != nil {
			logger.Fatal(err)
		}

		viper.SetConfigFile("../configs/secret_config.yaml")
		if err := viper.ReadInConfig(); err != nil {
			logger.Fatal(err)
		}
		instance.JwtKey = viper.GetString("jwt_auth_key")
	})

	return instance
}

func (config *Config) CorsSettings() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedMethods: config.CorseStorage.AllowedMethods,
		AllowedOrigins: config.CorseStorage.AllowedOrigins,
		AllowedHeaders: config.CorseStorage.AllowedHeader,
		ExposedHeaders: config.CorseStorage.ExposedHeaders,
	})
	return c
}
