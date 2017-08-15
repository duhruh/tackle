package tackle

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/duhruh/tackle/network/http"
)

type EndpointCollection interface {
	Get(name string) endpoint.Endpoint
	Add(name string, endpoint endpoint.Endpoint)
}

type endpointCollection struct {
	endpoints map[string]endpoint.Endpoint
}

func NewEndpointCollection() EndpointCollection {
	return endpointCollection{
		endpoints: make(map[string]endpoint.Endpoint),
	}
}

func (ec endpointCollection) Get(name string) endpoint.Endpoint {
	return ec.endpoints[name]
}

func (ec endpointCollection) Add(name string, endpoint endpoint.Endpoint) {
	ec.endpoints[name] = endpoint
}

type RegisterAdapter interface {
	Register(registry EndpointRegistry)
}

type EndpointRegistry interface {
	RegisterHttpEndpoint(he http.HttpEndpoint)
	HttpEndpoints() []http.HttpEndpoint
}

type endpointRegistry struct {
	httpEndpoints []http.HttpEndpoint
}

func NewEndpointRegistry() EndpointRegistry {
	return &endpointRegistry{}
}

func (er *endpointRegistry) RegisterHttpEndpoint(he http.HttpEndpoint) {
	er.httpEndpoints = append(er.httpEndpoints, he)
}

func (er *endpointRegistry) HttpEndpoints() []http.HttpEndpoint {
	return er.httpEndpoints
}

