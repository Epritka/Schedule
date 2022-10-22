package config

type Config struct {
	Port    int  `yaml:"Port"`
	IsDebug bool `yaml:"IsDebug"`
}
