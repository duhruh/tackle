package http

import "sync"

type AppHttpTransport interface {
	Build(transports []HttpTransport)
	Start(wg *sync.WaitGroup)
	Transports() []HttpTransport
}
