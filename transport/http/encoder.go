package http


import (
	"github.com/go-kit/kit/transport/http"
)

type Encoder interface {
	Encode() http.EncodeResponseFunc
	Decode() http.DecodeRequestFunc
}

type encoder struct{
	encode http.EncodeResponseFunc
	decode http.DecodeRequestFunc
}

func NewEncoder(decode http.DecodeRequestFunc, encode http.EncodeResponseFunc) Encoder{
	return encoder{
		decode: decode,
		encode: encode,
	}
}

func (e encoder) Encode() http.EncodeResponseFunc { return e.encode }
func (e encoder) Decode() http.DecodeRequestFunc { return e.decode }