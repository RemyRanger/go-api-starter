// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package ports

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Fetch Flows
	// (GET /flows)
	FindFlows(w http.ResponseWriter, r *http.Request, params FindFlowsParams)
	// Create Flow
	// (POST /flows)
	CreateFlow(w http.ResponseWriter, r *http.Request)
	// Fetch Flows origins
	// (GET /flows/origins)
	FindFlowsOrigins(w http.ResponseWriter, r *http.Request, params FindFlowsOriginsParams)
	// Fetch global Flows statistics
	// (GET /flows/stats)
	FindFlowsStatistics(w http.ResponseWriter, r *http.Request)
	// Returns flow by ID
	// (GET /flows/{id})
	FindFlowByID(w http.ResponseWriter, r *http.Request, id string)
	// Update Flow
	// (PUT /flows/{id})
	UpdateFlow(w http.ResponseWriter, r *http.Request, id string)
	// Fetch Contacts by flow ID
	// (GET /flows/{id}/contacts)
	FindContactsByFlowID(w http.ResponseWriter, r *http.Request, id string, params FindContactsByFlowIDParams)
	// Update Flow Visibility
	// (PUT /flows/{id}/visibility)
	UpdateFlowVisibility(w http.ResponseWriter, r *http.Request, id string)
	// Returns jobs by flow job name code
	// (GET /flows/{job_name}/jobs)
	FindJobsByFlowJobName(w http.ResponseWriter, r *http.Request, jobName string, params FindJobsByFlowJobNameParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// FindFlows operation middleware
func (siw *ServerInterfaceWrapper) FindFlows(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params FindFlowsParams

	// ------------- Optional query parameter "order_by" -------------

	err = runtime.BindQueryParameter("form", true, false, "order_by", r.URL.Query(), &params.OrderBy)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "order_by", Err: err})
		return
	}

	// ------------- Optional query parameter "page_size" -------------

	err = runtime.BindQueryParameter("form", true, false, "page_size", r.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page_size", Err: err})
		return
	}

	// ------------- Optional query parameter "page_token" -------------

	err = runtime.BindQueryParameter("form", true, false, "page_token", r.URL.Query(), &params.PageToken)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page_token", Err: err})
		return
	}

	// ------------- Optional query parameter "sort_dir_desc" -------------

	err = runtime.BindQueryParameter("form", true, false, "sort_dir_desc", r.URL.Query(), &params.SortDirDesc)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sort_dir_desc", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FindFlows(w, r, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateFlow operation middleware
func (siw *ServerInterfaceWrapper) CreateFlow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateFlow(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// FindFlowsOrigins operation middleware
func (siw *ServerInterfaceWrapper) FindFlowsOrigins(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params FindFlowsOriginsParams

	// ------------- Optional query parameter "type_offers" -------------

	err = runtime.BindQueryParameter("form", true, false, "type_offers", r.URL.Query(), &params.TypeOffers)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "type_offers", Err: err})
		return
	}

	// ------------- Optional query parameter "status" -------------

	err = runtime.BindQueryParameter("form", true, false, "status", r.URL.Query(), &params.Status)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "status", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FindFlowsOrigins(w, r, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// FindFlowsStatistics operation middleware
func (siw *ServerInterfaceWrapper) FindFlowsStatistics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FindFlowsStatistics(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// FindFlowByID operation middleware
func (siw *ServerInterfaceWrapper) FindFlowByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FindFlowByID(w, r, id)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateFlow operation middleware
func (siw *ServerInterfaceWrapper) UpdateFlow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateFlow(w, r, id)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// FindContactsByFlowID operation middleware
func (siw *ServerInterfaceWrapper) FindContactsByFlowID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params FindContactsByFlowIDParams

	// ------------- Optional query parameter "order_by" -------------

	err = runtime.BindQueryParameter("form", true, false, "order_by", r.URL.Query(), &params.OrderBy)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "order_by", Err: err})
		return
	}

	// ------------- Optional query parameter "page_size" -------------

	err = runtime.BindQueryParameter("form", true, false, "page_size", r.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page_size", Err: err})
		return
	}

	// ------------- Optional query parameter "page_token" -------------

	err = runtime.BindQueryParameter("form", true, false, "page_token", r.URL.Query(), &params.PageToken)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page_token", Err: err})
		return
	}

	// ------------- Optional query parameter "sort_dir_desc" -------------

	err = runtime.BindQueryParameter("form", true, false, "sort_dir_desc", r.URL.Query(), &params.SortDirDesc)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sort_dir_desc", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FindContactsByFlowID(w, r, id, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateFlowVisibility operation middleware
func (siw *ServerInterfaceWrapper) UpdateFlowVisibility(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateFlowVisibility(w, r, id)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// FindJobsByFlowJobName operation middleware
func (siw *ServerInterfaceWrapper) FindJobsByFlowJobName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "job_name" -------------
	var jobName string

	err = runtime.BindStyledParameterWithLocation("simple", false, "job_name", runtime.ParamLocationPath, chi.URLParam(r, "job_name"), &jobName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "job_name", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params FindJobsByFlowJobNameParams

	// ------------- Optional query parameter "page_size" -------------

	err = runtime.BindQueryParameter("form", true, false, "page_size", r.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page_size", Err: err})
		return
	}

	// ------------- Optional query parameter "page_token" -------------

	err = runtime.BindQueryParameter("form", true, false, "page_token", r.URL.Query(), &params.PageToken)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page_token", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FindJobsByFlowJobName(w, r, jobName, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshallingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshallingParamError) Error() string {
	return fmt.Sprintf("Error unmarshalling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshallingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/flows", wrapper.FindFlows)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/flows", wrapper.CreateFlow)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/flows/origins", wrapper.FindFlowsOrigins)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/flows/stats", wrapper.FindFlowsStatistics)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/flows/{id}", wrapper.FindFlowByID)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/flows/{id}", wrapper.UpdateFlow)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/flows/{id}/contacts", wrapper.FindContactsByFlowID)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/flows/{id}/visibility", wrapper.UpdateFlowVisibility)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/flows/{job_name}/jobs", wrapper.FindJobsByFlowJobName)
	})

	return r
}
