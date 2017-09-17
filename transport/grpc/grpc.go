package grpc

import (
	"google.golang.org/grpc"
	"github.com/go-kit/kit/endpoint"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

type GrpcTransport interface {
	NewHandler(g *grpc.Server)
}

func NewServer(end endpoint.Endpoint, encoder Encoder, options []kitgrpc.ServerOption) *kitgrpc.Server {
	return kitgrpc.NewServer(
		end,
		encoder.Decode(),
		encoder.Encode(),
		options...,
	)
}
