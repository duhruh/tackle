package network

import "github.com/duhruh/tackle"

type Response interface {
	tackle.Create
}


type response struct {
	data map[string]interface{}
}


func NewResponse() Response {
	return &response{}
}

func (r *response) Put(key string, value interface{}) *Response {
	r.data[key] = value
	return r
}
func (r *response) Get(key string) interface{} {
	return r.data[key]
}