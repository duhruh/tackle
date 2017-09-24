package generic

import (
	"errors"
	"reflect"
)

type DynamicCaller interface {
	Call(class interface{}, end string, args ...interface{}) (interface{}, error)
}

type dynamicCaller struct{}

func NewDynamicCaller() DynamicCaller {
	return dynamicCaller{}
}

func (d dynamicCaller) Call(class interface{}, end string, args ...interface{}) (interface{}, error) {
	e := reflect.ValueOf(class)
	m := e.MethodByName(end)

	if !m.IsValid() {
		return nil, errors.New("function not found")
	}

	var in []reflect.Value
	for _, arg := range args {
		in = append(in, reflect.ValueOf(arg))
	}

	out := m.Call(in)

	return out[0].Interface(), nil
}
