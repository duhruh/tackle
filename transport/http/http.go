package http

import (
	"net/http"
	//"github.com/duhruh/tackle"
	//http2 "github.com/duhruh/scaffold/app/hello/transport/http"
	http3 "github.com/go-kit/kit/transport/http"
	//"github.com/duhruh/tackle"
	//"golang.org/x/net/http2"
	"github.com/go-kit/kit/endpoint"
)

type HttpTransport interface {
	NewHandler(m *http.ServeMux) http.Handler
}



func NewServer(end endpoint.Endpoint, serializer Serializer, options []http3.ServerOption) *http3.Server{
	return http3.NewServer(
		end,
		serializer.Deserialize(),
		serializer.Serialize(),
		options...,
	)
}
