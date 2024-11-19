package control

import (
	"errors"
	"go-api-starter/internal/common/config"
	"go-api-starter/internal/common/db"
	"go-api-starter/internal/common/logger"
	"go-api-starter/internal/common/middleware"
	"go-api-starter/internal/common/server"
	"net/http"

	apis_handler "go-api-starter/internal/services/apis/adapters/handler"
	apis_repository "go-api-starter/internal/services/apis/adapters/repository"
	apis_core "go-api-starter/internal/services/apis/core"
	apis_ports "go-api-starter/internal/services/apis/ports"

	oas_handler "go-api-starter/internal/services/oas/adapters/handler"
	oas_repository "go-api-starter/internal/services/oas/adapters/repository"
	oas_core "go-api-starter/internal/services/oas/core"
	oas_ports "go-api-starter/internal/services/oas/ports"

	"github.com/danibix95/zeropino/middlewares/std"
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

	// init router
	router := mux.NewRouter()

	// Register Middlewares
	router.Use(otelmux.Middleware(app_name))
	router.Use(std.RequestLogger(&log.Logger, nil))
	router.Use(middleware.HealthCheckMiddleware)

	// Register Handlers
	subRouter := router.PathPrefix("/v1").Subrouter()
	apis_ports.HandlerFromMux(apis_handler.NewHandler(apis_core.New(apis_repository.New(gormDb))), subRouter)
	oas_ports.HandlerFromMux(oas_handler.NewHandler(oas_core.New(oas_repository.New(gormDb), apis_repository.New(gormDb))), subRouter)

	router.NotFoundHandler = router.NewRoute().HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		server.RespondError(w, errors.New("not found"), http.StatusNotFound)
	}).GetHandler() // to let 404 propagate to middlewares

	return gormDb, router
}
