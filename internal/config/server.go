package config

import "time"

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`

	Buffers  `yaml:"buffers"`
	Timeouts `yaml:"timeouts"`
}

type Buffers struct {
	Read  int `yaml:"read"`
	Write int `yaml:"write"`
}

type Timeouts struct {
	Read  time.Duration `yaml:"read"`
	Write time.Duration `yaml:"write"`
	Idle  time.Duration `yaml:"idle"`
}