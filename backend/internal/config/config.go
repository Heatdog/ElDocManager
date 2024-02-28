package config

import (
	"ElDocManager/pkg/logging"
	"sync"

	"github.com/rs/cors"
	"github.com/spf13/viper"
)

type Config struct {
	IsDebug        bool                 `yaml:"is_debug"`
	BackendStorage ListenBackend        `yaml:"listen_backend"`
	CorseStorage   CorsStorageConfig    `yaml:"cors_settings"`
	PostgreStorage PostgreStorageConfig `yaml:"postgre_settings"`
}

type ListenBackend struct {
	Type   string `yaml:"type"`
	BindIp string `yaml:"bind_ip"`
	Port   string `yaml:"port"`
}

type CorsStorageConfig struct {
	AllowedMethods []string `yaml:"allowed_methods"`
	AllowedOrigins []string `yaml:"allowed_origins"`
	AllowedHeader  []string `yaml:"allowed_headers"`
	ExposedHeaders []string `yaml:"exposed_headers"`
}

type PostgreStorageConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
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
