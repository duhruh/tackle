package http

//import (
//	"github.com/duhruh/tackle"
//)
//
//type HttpBridge interface {
//	HttpEndpoint() tackle.HttpEndpoint
//	Serializer() Serializer
//
//}
//
//
//func NewHttpBridge(httpEndpoint tackle.HttpEndpoint, serializer Serializer) HttpBridge {
//	return httpBridge{
//		httpEndpoint: httpEndpoint,
//		serializer: serializer,
//	}
//}
//
//type httpBridge struct {
//	httpEndpoint tackle.HttpEndpoint
//	serializer Serializer
//}
//
//func (hb httpBridge) HttpEndpoint() tackle.HttpEndpoint{
//	return hb.httpEndpoint
//}
//func (hb httpBridge) Serializer() Serializer{
//	return hb.serializer
//}
