package oas

//go:generate oapi-codegen --old-config-style -o ports/handler_types.gen.go -package=ports -include-tags=Oas -generate types ../../../doc/dist/control.yaml
//go:generate oapi-codegen --old-config-style -o ports/handler_api.gen.go -package=ports -include-tags=Oas -generate gorilla-server ../../../doc/dist/control.yaml