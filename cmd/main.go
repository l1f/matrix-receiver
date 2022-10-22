package main

import (
	"github.com/rs/zerolog"
	"github.com/urfave/cli/v2"
	"os"
)

const version = "DEV"

var logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
var commands cli.Commands

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func RegisterCommand(command *cli.Command) {
	commands = append(commands, command)
}

func main() {
	app := &cli.App{
		Name:     "Matrix Alertmanager Receiver",
		HelpName: "Receives messages from alert manager and forwards them to matrix",
		Version:  version,
		Commands: commands,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Aliases:  []string{"c"},
				Usage:    "The path to the config file.",
				EnvVars:  []string{"MAR_CONFIG_PATH"},
				Required: true,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}
}
