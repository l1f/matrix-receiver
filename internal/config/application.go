package config

type Application struct {
	Message `yaml:"message"`
	Server  `yaml:"server"`
}

func (a Application) validate() error {
	if err := a.Message.validate(); err != nil {
		return err
	}

	if err := a.Server.validate(); err != nil {
		return err
	}

	return nil
}
