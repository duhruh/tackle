package tackle

import (
	"github.com/duhruh/wiki/lib/tackle/transport/grpc"
	"github.com/duhruh/wiki/lib/tackle/transport/http"
)

type Application interface {
	Build()
	Start()
	HttpTransport() http.AppHttpTransport
	GrpcTransport() grpc.AppGrpcTransport
}

type Environment string

const (
	Development Environment = "development"
	Production  Environment = "production"
	Test        Environment = "test"
)
