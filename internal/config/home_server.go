package config

import (
	"fmt"
	"net/url"
)

type HomeServer struct {
	URL string `yaml:"URL"`
}

func (h HomeServer) validate() error {
	_, err := url.ParseRequestURI(fmt.Sprintf("https://%s", h.URL))
	if err != nil {
		return fmt.Errorf("error parsing home server url: %s", err)
	}

	return nil
}
