package http

import (
	"net/http"
	kitlog "github.com/go-kit/kit/log"
)

type HttpBuilder interface {
	Namespace() string
	Mount (endpoints []HttpEndpoint, log kitlog.Logger) http.Handler
}
