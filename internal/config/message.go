package config

import "errors"

var (
	ErrTTLTooLow = errors.New("ttl to low")
)

type Message struct {
	TTL int `yaml:"ttl"`
}

func (m Message) validate() error {
	if m.TTL < 1 {
		return ErrTTLTooLow
	}

	return nil
}
