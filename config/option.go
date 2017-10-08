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

type OptionMap interface {
	Get(string) interface{}
	All() []Option
	AddMap(OptionMap)
}

type optionMap struct {
	options []Option
}

func NewOptionMap(o []Option) OptionMap {
	return &optionMap{options: o}
}

func (o *optionMap) Get(op string) interface{} {
	for _, option := range o.options {
		if option.Name() == op {
			return option.Value()
		}
	}
	var oo Option
	return oo
}

func (o *optionMap) All() []Option {
	return o.options
}

func (o *optionMap) AddMap(om OptionMap) {
	o.options = append(o.options, om.All()...)
}

//type Value interface {
//	Value() interface{}
//}
