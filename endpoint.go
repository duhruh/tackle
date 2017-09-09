package tackle

import (
	"github.com/go-kit/kit/endpoint"
)

type EndpointFactory interface {
	Generate(end string) (endpoint.Endpoint, error)
	GenerateWithInstance(class interface{}, end string) (endpoint.Endpoint, error)
}

type endpointFactory struct {
	DynamicCaller
}

func NewEndpointFactory() EndpointFactory {
	return endpointFactory{
		DynamicCaller: NewDynamicCaller(),
	}
}

func (ef endpointFactory) Generate(end string) (endpoint.Endpoint, error) {
	return ef.GenerateWithInstance(ef, end)
}

func (ef endpointFactory) GenerateWithInstance(class interface{}, end string) (endpoint.Endpoint, error) {
	result, err := ef.Call(class, end)
	return result.(endpoint.Endpoint), err
}
