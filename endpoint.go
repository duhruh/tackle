package tackle

import (
	"github.com/go-kit/kit/endpoint"
)

// EndpointCaller -
type EndpointCaller func() endpoint.Endpoint

// EndpointMap -
type EndpointMap map[string]EndpointCaller

// EndpointFactory -
type EndpointFactory interface {
	Generate(end string) (endpoint.Endpoint, error)
	SetEndpoints(EndpointMap)
}
