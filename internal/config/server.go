package config

import (
	"errors"
	"time"
)

var (
	ErrServerHostNotSet = errors.New("server host not set")
	ErrServerPortTooLow = errors.New("server port too low")
)

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`

	Buffers  `yaml:"buffers"`
	Timeouts `yaml:"timeouts"`
}

func (s Server) validate() error {
	if s.Host == "" {
		return ErrServerHostNotSet
	}

	if s.Port <= 1024 {
		return ErrServerPortTooLow
	}

	return nil
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
