package grpc

import (
	"github.com/go-kit/kit/endpoint"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

type GrpcTransport interface {
	NewHandler(g *grpc.Server)
	Handlers() []Handler
}

func NewServer(end endpoint.Endpoint, encoder Encoder, options []kitgrpc.ServerOption) *kitgrpc.Server {
	return kitgrpc.NewServer(
		end,
		encoder.Decode(),
		encoder.Encode(),
		options...,
	)
}
