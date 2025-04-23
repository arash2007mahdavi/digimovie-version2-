package config

import (
	"os"

	"github.com/spf13/viper"
)

func GetConfig() *Config {
	path := getPath(os.Getenv("APP_ENV"))
	v := loadConfig(path, "yml")
	config := parseConfig(v)
	return config
}

func getPath(env string) string {
	if env == "docker" {
		return "../config/docker-config"
	} else if env == "production" {
		return "../config/production-config"
	} else {
		return "../config/development-config"
	}
}

func loadConfig(filename string, filetype string) *viper.Viper {
	v := viper.New()
	v.SetConfigName(filename)
	v.SetConfigType(filetype)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return v
}

func parseConfig(v *viper.Viper) *Config {
	config := &Config{}
	err := v.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return config
}