package flows

//go:generate oapi-codegen --old-config-style -o ports/api_types.gen.go -package=ports -include-tags=Flows -generate types ../../../openapi/backoffice_v1.yaml
//go:generate oapi-codegen --old-config-style -o ports/api.gen.go -package=ports -include-tags=Flows -generate chi-server ../../../openapi/backoffice_v1.yaml

//go:generate mockgen -destination=mocks/sql_repository.go -package=mocks gitlab.com/cdv-projects/go-apis/internal/services/flows SQL
