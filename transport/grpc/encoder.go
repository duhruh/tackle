package grpc


import (
	"github.com/duhruh/tackle"
	"github.com/go-kit/kit/transport/grpc"
)

type Encoder interface {
	Encode() grpc.EncodeResponseFunc
	Decode() grpc.DecodeRequestFunc
}

type encoder struct {
	encode grpc.EncodeResponseFunc
	decode grpc.DecodeRequestFunc
}

func NewEncoder(decode grpc.DecodeRequestFunc, encode grpc.EncodeResponseFunc) Encoder {
	return encoder{
		decode: decode,
		encode: encode,
	}
}

func (e encoder) Encode() grpc.EncodeResponseFunc { return e.encode }
func (e encoder) Decode() grpc.DecodeRequestFunc  { return e.decode }

type EncoderFactory interface {
	Generate(end string) (Encoder, error)
	GenerateWithInstance(class interface{}, end string) (Encoder, error)
}

type encoderFactory struct {
	tackle.DynamicCaller
}

func NewEncoderFactory() EncoderFactory {
	return encoderFactory{
		DynamicCaller: tackle.NewDynamicCaller(),
	}
}

func (ef encoderFactory) Generate(end string) (Encoder, error) {
	return ef.GenerateWithInstance(ef, end)
}

func (ef encoderFactory) GenerateWithInstance(class interface{}, end string) (Encoder, error) {
	result, err := ef.Call(class, end)
	return result.(Encoder), err
}
