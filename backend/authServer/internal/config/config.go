package config

import (
	"sync"

	logger "github.com/Heatdog/ElDocManager/backend/logger/app"
	"github.com/spf13/viper"
)

type Config struct {
	TokenKey     string            `yaml:"secret_key"`
	AuthStorage  AuthServerStorage `mapstructure:"auth_server_listen"`
	RedisStorage RedisStorage      `mapstructure:"listen_redis"`
}

type AuthServerStorage struct {
	Type   string `mapstructure:"type"`
	BindIp string `mapstructure:"bind_ip"`
	Port   string `mapstructure:"port"`
}

type RedisStorage struct {
	Type            string `mapstructure:"type"`
	BindIp          string `mapstructure:"bind_ip"`
	Port            string `mapstructure:"port"`
	Password        string `mapstructure:"password"`
	TokenExpiration int    `mapstructure:"token_exparation_days"`
}

var instance *Config
var once sync.Once

func GetConfig(logger *logger.Logger) *Config {
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
	})

	return instance
}
