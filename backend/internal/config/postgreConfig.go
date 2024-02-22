package config

type PostgreStorageConfig struct {
	Username   string
	Password   string
	Host       string
	Port       string
	Database   string
	MaxAttemps int
}
