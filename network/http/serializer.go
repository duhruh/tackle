package http

import(
	"github.com/go-kit/kit/transport/http"
)


type Serializer interface {
	Decode() http.DecodeRequestFunc
	Encode() http.EncodeResponseFunc
	Error() http.ErrorEncoder
}

type serializer struct {
	decode http.DecodeRequestFunc
	encode http.EncodeResponseFunc
	error http.ErrorEncoder
}


func NewSerializer(decode http.DecodeRequestFunc, encode http.EncodeResponseFunc, error http.ErrorEncoder) Serializer{
	return serializer{decode:decode,encode:encode,error:error}
}


func (s serializer) Decode() http.DecodeRequestFunc {
	return s.decode
}
func (s serializer) Encode() http.EncodeResponseFunc {
	return s.encode
}
func (s serializer) Error() http.ErrorEncoder {
	return s.error
}