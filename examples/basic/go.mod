module github.com/vessel-api/vesselapi-go/examples/basic

go 1.24.0

require github.com/vessel-api/vesselapi-go/v4 v4.0.0

replace github.com/vessel-api/vesselapi-go/v4 => ../..

require (
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/oapi-codegen/runtime v1.6.0 // indirect
)
