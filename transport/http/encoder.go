package http

import (
	"github.com/duhruh/tackle"
	"github.com/go-kit/kit/transport/http"
)

type Encoder interface {
	Encode() http.EncodeResponseFunc
	Decode() http.DecodeRequestFunc
}

type encoder struct {
	encode http.EncodeResponseFunc
	decode http.DecodeRequestFunc
}

func NewEncoder(decode http.DecodeRequestFunc, encode http.EncodeResponseFunc) Encoder {
	return encoder{
		decode: decode,
		encode: encode,
	}
}

func (e encoder) Encode() http.EncodeResponseFunc { return e.encode }
func (e encoder) Decode() http.DecodeRequestFunc  { return e.decode }

type EncoderFactory interface {
	Generate(end string) (Encoder, error)
	GenerateWithInstance(class interface{}, end string) (Encoder, error)
	ErrorEncoder() http.ErrorEncoder
}

type encoderFactory struct {
	tackle.DynamicCaller
}

func NewEncoderFactory() EncoderFactory {
	return encoderFactory{
		DynamicCaller: tackle.NewDynamicCaller(),
	}
}

func (ef encoderFactory) ErrorEncoder() http.ErrorEncoder { return nil }

func (ef encoderFactory) Generate(end string) (Encoder, error) {
	return ef.GenerateWithInstance(ef, end)
}

func (ef encoderFactory) GenerateWithInstance(class interface{}, end string) (Encoder, error) {
	result, err := ef.Call(class, end)
	return result.(Encoder), err
}
