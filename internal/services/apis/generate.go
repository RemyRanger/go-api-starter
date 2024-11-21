package apis

//go:generate oapi-codegen --old-config-style -o ports/handler_types.gen.go -package=ports -include-tags=Api -generate types ../../../doc/dist/control.yaml
//go:generate oapi-codegen --old-config-style -o ports/handler_api.gen.go -package=ports -include-tags=Api -generate gorilla-server ../../../doc/dist/control.yaml

//go:generate mockgen -destination=mocks/sql_repository.go -package=mocks go-api-starter/internal/services/apis/ports Repository
