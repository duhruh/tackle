package config

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
