package config

import (
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type YamlLoader interface {
	Loader
}

type yamlLoader struct {
}

func NewYamlLoader() YamlLoader {
	return yamlLoader{}
}

func (y yamlLoader) LoadFromFile(reader io.Reader) (Config, error) {
	var config Config
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return config, err
	}

	options := y.loadYamlFromData([]byte(data))
	config = NewConfig(options)

	return config, nil
}

func (y yamlLoader) loadYamlFromData(data []byte) OptionMap {
	m := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		panic(err)
	}

	return y.resolveMap(m)
}
func (y yamlLoader) resolveMap(m map[interface{}]interface{}) OptionMap {
	var uhh []Option

	for k, v := range m {

		if k.(string) == "include>" {
			om := y.loadFile(v.(string))
			uhh = append(uhh, om.All()...)
			continue
		}

		o := NewOption(k.(string))
		o.SetValue(y.resolveType(v))

		uhh = append(uhh, o)
	}

	return NewOptionMap(uhh)
}

func (y yamlLoader) loadFile(f string) OptionMap {
	ff, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(ff)
	if err != nil {
		panic(err)
	}

	return y.loadYamlFromData([]byte(data))
}

func (y yamlLoader) resolveType(v interface{}) interface{} {
	switch v.(type) {
	case []interface{}:
		return y.resolveArrayThings(v.([]interface{}))
	case map[interface{}]interface{}:
		return y.resolveMap(v.(map[interface{}]interface{}))
	default:
		return v
	}
}

func (y yamlLoader) resolveArrayThings(v []interface{}) []interface{} {
	var arr []interface{}
	for _, uh := range v {
		arr = append(arr, y.resolveType(uh))
	}
	return arr
}
