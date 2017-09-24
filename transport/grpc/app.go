package grpc

import "sync"

type AppGrpcTransport interface {
	Build(transports []GrpcTransport)
	Start(wg *sync.WaitGroup)
	Transports() []GrpcTransport
}
