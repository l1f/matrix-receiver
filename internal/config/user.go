package config

import (
	"errors"
	"regexp"
)

var usernameRegex = regexp.MustCompile("^@[a-zA-Z0-9_]+:([a-zA-Z0-9-_]+.)+[a-zA-Z0-9-_]+$")
var (
	ErrPasswordAndTokenEmpty = errors.New("there must be at least one password or token present")
	ErrInvalidUsernameFormat = errors.New("invalid username format")
	ErrHomeServerIsEmpty     = errors.New("home server is empty")
)

type User struct {
	HomeServer string `yaml:"home_server"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password,omitempty"`
	Token      string `yaml:"token,omitempty"`
}

func (u User) validate() error {
	if u.Password == "" && u.Token == "" {
		return ErrPasswordAndTokenEmpty
	}

	if !usernameRegex.MatchString(u.Username) {
		return ErrInvalidUsernameFormat
	}

	if u.HomeServer == "" {
		return ErrHomeServerIsEmpty
	}

	return nil
}
