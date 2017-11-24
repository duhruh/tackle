package task

type Argument interface {
	CommandLineArgument
}

type argument struct {
	key          string
	defaultValue string
	value        interface{}
	description  string
}

func NewArgument(key string, description string, ops ...interface{}) Argument {
	return &argument{
		key:         key,
		description: description,
	}
}

func (o *argument) SetValue(v interface{}) {
	o.value = v
}
func (o *argument) Key() string {
	return o.key
}
func (o *argument) DefaultValue() interface{} {
	return o.defaultValue
}
func (o *argument) Value() interface{} {
	return o.value
}
func (o *argument) Description() string {
	return o.description
}
