package config

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Config structure to hold the configuration values
type Config struct {
	Server    Server    `mapstructure:"server" validate:"required"`
	Telemetry Telemetry `mapstructure:"telemetry" validate:"required"`
	Db        Db        `mapstructure:"db" validate:"required"`
	Logs      Logs      `mapstructure:"logs" validate:"required"`
}

type Server struct {
	Env  string `mapstructure:"env" validate:"required,oneof=DEV PROD"`
	Addr string `mapstructure:"addr" validate:"required"`
}

type Telemetry struct {
	Addr string `mapstructure:"addr" validate:"required"`
}

type Db struct {
	Addr string `mapstructure:"addr" validate:"required"`
}

type Logs struct {
	Level string `mapstructure:"level" validate:"required,oneof=INFO DEBUG TRACE"`
}

func LoadConfig(configName string) (Config, error) {
	// Set the name of the config file (without extension)
	viper.SetConfigName(configName)
	// Set the path to look for the config file
	viper.AddConfigPath("./config/")     // config folder from current directory
	viper.AddConfigPath("../../config/") // config folder from current directory
	// Set the file type to YAML
	viper.SetConfigType("yaml")

	// Automatic environment variables
	viper.AutomaticEnv()

	// Replace `.` in env variables with `_` (ex: env var DATABASE_HOST is replaced by DATABASE.HOST to match struct config )
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Attempt to read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		// If no configuration file is found, fallback to env variables
		if viper_err, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn().Msgf("%s, falling back to environment variables", viper_err.Error())
		} else {
			// Handle other errors
			return Config{}, fmt.Errorf("error reading config file: %w", err)
		}
	} else {
		log.Info().Msg("Config file loaded successfully")
	}

	// Unmarshal the configuration into the struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, fmt.Errorf("unable to decode config: %w", err)
	}

	// Validate the config
	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		// Return validation errors
		return Config{}, fmt.Errorf("unable to validate config: %w", err)
	}

	return config, nil
}
