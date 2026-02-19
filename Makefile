.PHONY: all build test vet fmt tidy generate smoke

all: fmt tidy build vet test

build:
	go build ./...

test:
	go test -race -v ./...

vet:
	go vet ./...

fmt:
	gofmt -s -w .

tidy:
	go mod tidy

generate:
	oapi-codegen --generate types,client --package vesselapi -o generated.go openapi/openapi.json

smoke:
	go test -race -v -tags=smoke -timeout 300s ./...
