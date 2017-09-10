package http

import (
	"github.com/go-kit/kit/transport/http"
)

type Serializer interface {
	Serialize() http.EncodeResponseFunc
	Deserialize() http.DecodeRequestFunc
}

type serializer struct {
	serialize   http.EncodeResponseFunc
	deserialize http.DecodeRequestFunc
}

func NewSerializer(deserialize http.DecodeRequestFunc, serialize http.EncodeResponseFunc) Serializer {
	return serializer{
		serialize:   serialize,
		deserialize: deserialize,
	}
}

func (s serializer) Serialize() http.EncodeResponseFunc  { return s.serialize }
func (s serializer) Deserialize() http.DecodeRequestFunc { return s.deserialize }
