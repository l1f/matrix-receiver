package main

import (
	"github.com/urfave/cli/v2"

  "matrix-alertmanager/cmd/receiver"
	"matrix-alertmanager/internal/config"
)

func init() {
	var command = cli.Command{
		Name:    "run",
		Usage:   "Starts the matrix-alertmanager service",
		Aliases: []string{"start"},
		Action:  run,
	}

	RegisterCommand(&command)
}

func run(ctx *cli.Context) error {
	logger.Info().Msg("Stating receiver..")

	path := ctx.String("config")
	c, err := config.Load(path)
	if err != nil {
		return err
	}

	return receiver.New(*c, logger)
}
