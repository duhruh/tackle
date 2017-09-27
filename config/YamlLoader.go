package config

import (
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
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

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(data), &m)

	println(m)
	if err != nil {
		return config, err
	}

}
