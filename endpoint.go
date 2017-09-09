package tackle

import (
	"errors"
	"github.com/go-kit/kit/endpoint"
	"reflect"
)

type EndpointFactory interface {
	Generate(end string) (endpoint.Endpoint, error)
	GenerateWithInstance(class interface{}, end string) (endpoint.Endpoint, error)
}

type endpointFactory struct {
}

func NewEndpointFactory() EndpointFactory {
	return endpointFactory{}
}

func (ef endpointFactory) Generate(end string) (endpoint.Endpoint, error) {
	return ef.GenerateWithInstance(ef, end)
}

func (ef endpointFactory) GenerateWithInstance(class interface{}, end string) (endpoint.Endpoint, error) {
	var factoryPoint endpoint.Endpoint

	e := reflect.ValueOf(class)
	m := e.MethodByName(end)

	if !m.IsValid() {
		return factoryPoint, errors.New("endpoint not found")
	}

	var in []reflect.Value
	out := m.Call(in)

	factoryPoint = out[0].Interface().(endpoint.Endpoint)

	return factoryPoint, nil
}
