package task

type Option interface {
	CommandLineArgument
}

type option struct {
	key          string
	defaultValue string
	value        interface{}
	description  string
}

func NewOption(key string, description string, ops ...interface{}) Option {
	return &option{
		key:         key,
		description: description,
	}
}

func (o *option) Key() string {
	return o.key
}
func (o *option) DefaultValue() interface{} {
	return o.defaultValue
}
func (o *option) Value() interface{} {
	return o.value
}
func (o *option) Description() string {
	return o.description
}
func (o *option) SetValue(v interface{}) {
	o.value = v
}
