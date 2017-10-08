package config

import (
	//"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	//"reflect"
)

type YamlLoader interface {
}

type yamlLoader struct {
}

func NewYamlLoader() Loader {
	return yamlLoader{}
}

func (y yamlLoader) LoadFromFile(reader io.Reader) (Config, error) {
	var config Config
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return config, err
	}

	options := loadYamlFromData([]byte(data))
	config = NewConfig(options)

	return config, nil
}

func loadYamlFromData(data []byte) OptionMap {
	m := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		panic(err)
	}

	return resolveMap(m)
}
func resolveMap(m map[interface{}]interface{}) OptionMap {
	var uhh []Option

	//println("=======")
	for k, v := range m {

		if k.(string) == "include>" {
			om := loadFile(v.(string))
			uhh = append(uhh, om.All()...)
			continue
		}
		//println(k.(string))
		o := NewOption(k.(string))
		o.SetValue(resolveType(v))

		uhh = append(uhh, o)
	}

	return NewOptionMap(uhh)
}

func loadFile(f string) OptionMap {
	ff, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(ff)
	if err != nil {
		panic(err)
	}

	return loadYamlFromData([]byte(data))
}

func resolveType(v interface{}) interface{} {
	//fmt.Println(reflect.TypeOf(v))
	switch v.(type) {
	case []interface{}:
		return resolveArrayThings(v.([]interface{}))
	case map[interface{}]interface{}:
		return resolveMap(v.(map[interface{}]interface{}))
	default:
		return v
	}
}

func resolveArrayThings(v []interface{}) []interface{} {
	var arr []interface{}
	for _, uh := range v {
		arr = append(arr, resolveType(uh))
	}
	return arr
}
