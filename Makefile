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
# Drop the generator's package comment. go/doc concatenates package comments in
# filename order, so leaving it here would make it the package overview on
# pkg.go.dev instead of the one in vesselapi.go.
	sed -i '/^package vesselapi$$/,$$!d' generated.go
	sed -i 's|https://api\.deepmap\.com/dev-test|https://api.vesselapi.com/v1|g; s|https://api\.deepmap\.com|https://api.vesselapi.com/v1|g' generated.go
	gofmt -w generated.go

smoke:
	go test -race -v -tags=smoke -timeout 300s ./...
