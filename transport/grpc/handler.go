package grpc

type Handler interface {
	Name() string
	Endpoint() string
	Encoder() string
}

type handler struct {
	name     string
	endpoint string
	encoder  string
}

func (r handler) Name() string     { return r.name }
func (r handler) Endpoint() string { return r.endpoint }
func (r handler) Encoder() string  { return r.encoder }

func NewHandler(name string, endpoint string, encoder string) Handler {
	return handler{
		name:     name,
		endpoint: endpoint,
		encoder:  encoder,
	}
}
