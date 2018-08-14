package http

import (
	"context"
	"net/http"

	"github.com/duhruh/wiki/lib/tackle/transport"
	"github.com/go-kit/kit/endpoint"
	http3 "github.com/go-kit/kit/transport/http"
)

// type TackleResponse func(ctx context.Context, w http.ResponseWriter, response Packet) error

type HttpTransport interface {
	NewHandler(m *http.ServeMux, options []http3.ServerOption) http.Handler
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

// TackleResponse  -
type TackleResponse func(ctx context.Context, w http.ResponseWriter, response transport.Packet) error
