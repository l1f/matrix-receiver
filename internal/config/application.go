package config

type Application struct {
	Server  `yaml:"server"`
}

func (a Application) validate() error {
	if err := a.Server.validate(); err != nil {
		return err
	}

	return nil
}
