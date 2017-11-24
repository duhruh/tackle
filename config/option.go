package config

type Option interface {
	Name() string
	Value() interface{}
	SetValue(interface{})
	SetName(string)
}

type option struct {
	name  string
	value interface{}
}

func (o *option) Name() string {
	return o.name
}

func (o *option) Value() interface{} {
	return o.value
}

func (o *option) SetValue(v interface{}) {
	o.value = v
}
func (o *option) SetName(n string) {
	o.name = n
}
func NewOption(n string) Option {
	return &option{name: n}
}
