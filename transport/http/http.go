package http

import (
	"github.com/go-kit/kit/endpoint"
	http3 "github.com/go-kit/kit/transport/http"
	"net/http"
)

type HttpTransport interface {
	NewHandler(m *http.ServeMux) http.Handler
	Routes() []Route
}

func NewServer(end endpoint.Endpoint, encoder Encoder, options []http3.ServerOption) *http3.Server {
	return http3.NewServer(
		end,
		encoder.Decode(),
		encoder.Encode(),
		options...,
	)
}
