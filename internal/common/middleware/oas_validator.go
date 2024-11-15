package middleware

import (
	"context"
	"errors"
	"fmt"
	"go-api-starter/internal/common/models"
	"net/http"
	"strings"

	"github.com/allegro/bigcache/v3"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/uptrace/go-clickhouse/ch"
	"go.opentelemetry.io/otel/trace"
)

// OapiRequestValidatorWithOptions Creates middleware to validate request by openapi specification.
func OapiRequestValidator(db *ch.DB, cache *bigcache.BigCache, options *openapi3filter.Options) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			span := trace.SpanFromContext(r.Context())
			defer span.End()

			// Remove path prefix
			splited_path := strings.Split(r.URL.Path, "/")
			swagger, err := fetchOASfromCacheOrDb(r.Context(), db, cache, splited_path[1])
			if err != nil {
				panic(err)
			}
			r.URL.Path = strings.ReplaceAll(r.URL.Path, "/"+splited_path[1], "")

			// WARN: OapiRequestValidatorWithOptions called with an OpenAPI spec that has `Servers` set. This may lead to an HTTP 400 with `no matching operation was found` when sending a valid request, as the validator performs `Host` header validation. If you're expecting `Host` header validation, you can silence this warning by setting `Options.SilenceServersWarning = true`. See https://github.com/deepmap/oapi-codegen/issues/882 for more information.
			swagger.Servers = nil

			// Create validator router
			router, err := gorillamux.NewRouter(swagger)
			if err != nil {
				panic(err)
			}

			// validate request
			if statusCode, err := validateRequest(r, router, options); err != nil {
				http.Error(w, err.Error(), statusCode)
				return
			}

			// serve
			next.ServeHTTP(w, r)
		})
	}
}

// validateRequest is called from the middleware above and actually does the work
// of validating a request.
func validateRequest(r *http.Request, router routers.Router, options *openapi3filter.Options) (int, error) {
	// Find route
	route, pathParams, err := router.FindRoute(r)
	if err != nil {
		return http.StatusNotFound, err // We failed to find a matching route for the request.
	}

	// Validate request
	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request:    r,
		PathParams: pathParams,
		Route:      route,
	}

	if options != nil {
		requestValidationInput.Options = options
	}

	if err := openapi3filter.ValidateRequest(r.Context(), requestValidationInput); err != nil {
		me := openapi3.MultiError{}
		if errors.As(err, &me) {
			return http.StatusBadRequest, me
		}

		switch e := err.(type) {
		case *openapi3filter.RequestError:
			// We've got a bad request
			// Split up the verbose error by lines and return the first one
			// openapi errors seem to be multi-line with a decent message on the first
			errorLines := strings.Split(e.Error(), "\n")
			return http.StatusBadRequest, fmt.Errorf(errorLines[0])
		/* case *openapi3filter.SecurityRequirementsError:
		return http.StatusUnauthorized, err */
		default:
			// This should never happen today, but if our upstream code changes,
			// we don't want to crash the server, so handle the unexpected error.
			return http.StatusInternalServerError, fmt.Errorf("error validating route: %s", err.Error())
		}
	}

	return http.StatusOK, nil
}

// fetchOASfromCacheOrDb return corresponding oas file from cache or db
func fetchOASfromCacheOrDb(ctx context.Context, db *ch.DB, cache *bigcache.BigCache, oasId string) (*openapi3.T, error) {
	// Try to find oas in cache
	if entry, err := cache.Get(oasId); err == nil {
		loader := &openapi3.Loader{Context: ctx}
		doc, err := loader.LoadFromData(entry)
		if err != nil {
			return nil, err
		}

		return doc, nil
	}

	oas := new(models.Oas)
	if err := db.NewSelect().Model(oas).Where("uuid = ?", oasId).Scan(ctx); err != nil {
		return nil, err
	}

	binaryOas, err := oas.JsonSpec.Value()
	if err != nil {
		return nil, err
	}

	loader := &openapi3.Loader{Context: ctx}
	doc, err := loader.LoadFromData(binaryOas.([]byte))
	if err != nil {
		return nil, err
	}

	// save to cache
	cache.Set(oasId, binaryOas.([]byte))

	return doc, nil
}
