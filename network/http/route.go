package http



type Route interface {
	Path() string
	Method() string
	Namespace() string
}


type route struct {
	path string
	method string
	namespace string
}



func NewRoute(path string, method string, namespace string) Route{
	return route{path, method, namespace}
}


func (r route) Path() string {
	return r.path
}

func (r route) Method() string {
	return r.method
}

func (r route) Namespace() string {
	return r.namespace
}
