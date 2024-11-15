package main

import (
	"context"
	"go-api-starter/internal/common/cache"
	"go-api-starter/internal/common/config"
	"go-api-starter/internal/common/db"
	"go-api-starter/internal/common/logger"
	"go-api-starter/internal/common/middleware"
	"go-api-starter/internal/common/server"
	"net/http"

	gateway_handler "go-api-starter/internal/services/gateway/adapters/handler"

	"github.com/danibix95/zeropino/middlewares/std"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

const app_name = "gateway"
const app_config_filename = "gateway-config"

func main() {
	// Init config
	app_config, err := config.LoadConfig(app_config_filename)
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading config")
	}

	// Init Logger
	logger.NewZerolog(app_config)

	// Run server
	if err := server.Run(app_name, app_config, newHTTPHandler(app_config)); err != nil {
		log.Fatal().Err(err)
	}
}

func newHTTPHandler(app_config config.Config) http.Handler {
	// init db
	db := db.NewClickHouse(app_config)

	// init cache
	cache := cache.New()

	// init router
	router := mux.NewRouter()

	// Register middlewares
	router.Use(std.RequestLogger(&log.Logger, nil))
	router.Use(otelmux.Middleware(app_name))
	router.Use(middleware.OapiRequestValidator(db, cache, &openapi3filter.Options{
		// We set a fake AuthenticationFunc to validate oas with Security section without errors
		AuthenticationFunc: func(c context.Context, input *openapi3filter.AuthenticationInput) error {
			return nil
		},
	}))

	// Register Handlers
	router.PathPrefix("/{key}").HandlerFunc(gateway_handler.NewHandler().HandleGateway)

	//todo
	/* proxy := httputil.NewSingleHostReverseProxy(tgt) */

	return router
}
