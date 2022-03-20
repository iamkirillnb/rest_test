
gen:
	go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
	oapi-codegen ./internal/oapi/law/openapi.yaml > ./internal/oapi/law/openapi.gen.go