package http

import "github.com/go-kit/kit/endpoint"

//
//import "github.com/go-kit/kit/endpoint"
//
type HttpEndpoint interface {
	Endpoint() endpoint.Endpoint
	Route() Route
	Serializer() Serializer
}

type httpEndpoint struct {
	endpoint endpoint.Endpoint
	route Route
	serializer Serializer
}

func NewHttpEndpoint(endpoint endpoint.Endpoint, route Route, serializer Serializer) HttpEndpoint {
	return httpEndpoint{
		endpoint: endpoint,
		route: route,
		serializer: serializer,
	}
}

func (he httpEndpoint) Endpoint() endpoint.Endpoint {
	return he.endpoint
}

func (he httpEndpoint) Route() Route {
	return he.route
}

func (he httpEndpoint) Serializer() Serializer{
	return hb.serializer
}
