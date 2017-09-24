package http

type Route interface {
	Method() string
	Path() string
	Endpoint() string
	Encoder() string
}

type route struct {
	method   string
	path     string
	endpoint string
	encoder  string
}

func (r route) Method() string   { return r.method }
func (r route) Path() string     { return r.path }
func (r route) Endpoint() string { return r.endpoint }
func (r route) Encoder() string  { return r.encoder }

func NewRoute(method string, path string, endpoint string, encoder string) Route {
	return route{
		method:   method,
		path:     path,
		endpoint: endpoint,
		encoder:  encoder,
	}
}
