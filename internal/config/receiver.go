package config

import (
	"errors"
)

var (
	ErrUsernameEmpty = errors.New("username empty")
	ErrRoomEmpty     = errors.New("room empty")
)

type Receiver struct {
	User string `yaml:"user"`
	Room string `yaml:"room"`
}

func (r Receiver) validate() error {
	if r.User == "" {
		return ErrUsernameEmpty
	}

	if r.Room == "" {
		return ErrRoomEmpty
	}

	return nil
}
