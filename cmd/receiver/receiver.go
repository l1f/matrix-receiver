package receiver

import (
	"github.com/rs/zerolog"
	"matrix-alertmanager/internal/application"
	"matrix-alertmanager/internal/config"
	"matrix-alertmanager/internal/logic"
	"matrix-alertmanager/internal/queue"
	"matrix-alertmanager/internal/server"
)

func New(config config.Config, logger zerolog.Logger) error {
	var ctx = &application.Context{
		Config:  config,
		Logger:  logger,
		Clients: nil,
	}

	clients, err := createClients(*ctx)
	if err != nil {
		return err
	}

	ctx.Clients = clients

	defaultQueue := queue.New(ctx.Config.Message.TTL, ctx.Logger)
	ctx.Queue = defaultQueue

	go ctx.Queue.Run()

	defaultLogic := logic.New(ctx)
	ctx.Logic = defaultLogic

	mainServer, listener, err := server.New(*ctx)
	if err != nil {
		return err
	}

	if err = mainServer.Serve(listener); err != nil {
		return err
	}

	return nil
}
