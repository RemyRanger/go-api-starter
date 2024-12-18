package main

import (
	"go-api-starter/internal/app/gateway"
	"go-api-starter/internal/common/config"
	"go-api-starter/internal/common/server"
	"go-api-starter/internal/common/telemetry"

	"github.com/rs/zerolog/log"
)

const app_name = "gateway"
const app_config_filename = "gateway-config"

func main() {
	// Init openTelemetry.
	otelShutdown := telemetry.SetupOTelSDK(app_name)

	// Init config
	app_config := config.LoadConfig(app_config_filename)

	// Run server
	_, handler := gateway.NewApp(app_name, app_config)
	if err := server.Run(app_name, app_config, handler, otelShutdown); err != nil {
		log.Fatal().Err(err).Msg("Error starting server")
	}
}
