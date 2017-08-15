package network

import (
	"github.com/duhruh/tackle"
	"github.com/duhruh/tackle/network/http"
)

type Transport interface {
	Start()
}

type HttpTransport interface {
	Transport
	Build(builders []http.HttpBuilder, endpoints []tackle.HttpEndpoint) HttpTransport
}
