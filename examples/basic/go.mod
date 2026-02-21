module github.com/vessel-api/vesselapi-go/examples/basic

go 1.22

require github.com/vessel-api/vesselapi-go/v3 v3.0.0

replace github.com/vessel-api/vesselapi-go/v3 => ../..

require (
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/oapi-codegen/runtime v1.1.2 // indirect
)
