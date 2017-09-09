package http

import (
	"net/http"
	//"github.com/duhruh/tackle"
	//http2 "github.com/duhruh/scaffold/app/hello/transport/http"
	//http3 "github.com/go-kit/kit/transport/http"
)

type HttpTransport interface {
	NewHandler(m *http.ServeMux) http.Handler
}



//func NewServer(end string, ef tackle.EndpointFactory, serializer http2.HttpSerializer, options ...http3.ServerOption){
//	callPoint := end + "Endpoint"
//	deserializer := end + "Request"
//	serialize := end + "Response"
//}
