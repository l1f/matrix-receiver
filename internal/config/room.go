package config

import "errors"

var (
	ErrRoomIdAndAddressNotSet = errors.New("room id and room address is not set")
)

type Room struct {
	Id      string `yaml:"id"`
	Address string `yaml:"address"`
}

func (r Room) validate() error {
	if r.Id == "" && r.Address == "" {
		return ErrRoomIdAndAddressNotSet
	}

	return nil
}