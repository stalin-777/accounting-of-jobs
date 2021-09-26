package config

import (
	"os"

	"github.com/juju/errors"
	"github.com/spf13/viper"
)

type Config struct {
	WebPort int `validate:"required"`
	DB      struct {
		Host     string `validate:"required"`
		Port     uint16 `validate:"required"`
		Username string `validate:"required"`
		Password string `validate:"required"`
		Dbname   string `validate:"required"`
	}
	Logger struct {
		FileName       string `validate:"required"`
		Path           string `validate:"required"`
		MaxSize        int    `validate:"required"`
		MaxRequestSize int    `validate:"required"`
		MaxBackups     int    `validate:"required"`
		MaxAge         int    `validate:"required"`
	}
}

func Init() (*Config, error) {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	cfg := &Config{}

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Annotate(err, "Failed to load configuration file")
	}

	if _, err := os.Stat("config.local.yaml"); err == nil {

		viper.SetConfigName("config.local")

		if err := viper.MergeInConfig(); err != nil {
			return nil, errors.Annotate(err, "Failed to merge configuration file")
		}
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Annotate(err, "Failed to unmarshal configuration file")
	}

	return cfg, nil
}
