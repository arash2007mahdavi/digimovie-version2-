package config

import "time"

type Config struct {
	Server   serverConfig
	Postgres postgresConfig
	Logger   loggerConfig
	Redis    redisConfig
}

type redisConfig struct {
	Port     int
	Host     string
	Password string
}
type loggerConfig struct {
	FilePath string
	Encoding string
	Level    string
	Logger   string
}

type serverConfig struct {
	Port int
	Host string
}

type postgresConfig struct {
	Host            string
	User            string
	Password        string
	Dbname          string
	Port            int
	Sslmode         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime *time.Duration
}
