package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port    int  `yaml:"Port"`
	IsDebug bool `yaml:"IsDebug"`
}

func New(filepath string) (Config, error) {
	var config Config

	yamlFile, err := ioutil.ReadFile(filepath)

	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
