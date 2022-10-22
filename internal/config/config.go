package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Application `yaml:"application"`

	HomeServers map[string]HomeServer `yaml:"home_servers"`
	Users       map[string]User       `yaml:"users"`
	Rooms       map[string]Room       `yaml:"rooms"`
	Receiver    map[string]Receiver   `yaml:"receiver"`
}

func (c Config) Validate() error {
	for _, homeServer := range c.HomeServers {
		if err := homeServer.validate(); err != nil {
			return err
		}
	}

	for _, user := range c.Users {
		if err := user.validate(); err != nil {
			return err
		}
	}

	for _, room := range c.Rooms {
		if err := room.validate(); err != nil {
			return err
		}
	}

	for _, receiver := range c.Receiver {
		if err := receiver.validate(); err != nil {
			return err
		}
	}

	return nil
}

func Load(path string) (*Config, error) {
	var config Config
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	if err = config.Validate(); err != nil {
		return nil, err
	}

	return &config, nil
}
