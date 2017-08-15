package network

import "github.com/duhruh/tackle"

type Request interface {
	tackle.Create
}


type request struct {
	data map[string]interface{}
}


func NewRequest() Request {
	return &request{}
}

func (r *request) Put(key string, value interface{}) Request{
	r.data[key] = value
	return r
}
func (r *request) Get(key string) interface{} {
	return r.data[key]
}