package tackle

import "github.com/duhruh/tackle/network/http"

type HttpApp interface {
	http.HttpBuilder
	RegisterAdapter
}