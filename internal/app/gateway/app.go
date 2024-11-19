package gateway

import (
	"context"
	"go-api-starter/internal/common/cache"
	"go-api-starter/internal/common/config"
	"go-api-starter/internal/common/db"
	"go-api-starter/internal/common/logger"
	"go-api-starter/internal/common/middleware"

	gateway_handler "go-api-starter/internal/services/gateway/adapters/handler"

	"github.com/danibix95/zeropino/middlewares/std"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"gorm.io/gorm"
)

func NewApp(app_name string, app_config config.Config) (*gorm.DB, *mux.Router) {
	// Init logger
	logger.NewZerolog(app_config)

	// Init db
	gormDb := db.NewPostgres(app_config)

	// Init Cache
	cache := cache.New()

	// init router
	router := mux.NewRouter()

	// Register middlewares
	router.Use(std.RequestLogger(&log.Logger, nil))
	router.Use(otelmux.Middleware(app_name))
	router.Use(middleware.OapiRequestValidator(gormDb, cache, &openapi3filter.Options{
		// We set a fake AuthenticationFunc to validate oas with Security section without errors
		AuthenticationFunc: func(c context.Context, input *openapi3filter.AuthenticationInput) error {
			return nil
		},
	}))

	// Register Handlers
	router.PathPrefix("/{key}").HandlerFunc(gateway_handler.NewHandler().HandleGateway)

	//todo
	/* proxy := httputil.NewSingleHostReverseProxy(tgt) */

	return gormDb, router
}
