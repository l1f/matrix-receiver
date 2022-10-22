package server

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"matrix-alertmanager/internal/application"
	"net"
	"strconv"
)

func New(ctx application.Context) (server *fasthttp.Server, listener net.Listener, err error) {
	config := ctx.Config

	server = &fasthttp.Server{
		ErrorHandler:    handleError(),
		Handler:         requestRouter(ctx),
		ReadBufferSize:  config.Server.Buffers.Read,
		WriteBufferSize: config.Server.Buffers.Write,
		ReadTimeout:     config.Server.Timeouts.Read,
		WriteTimeout:    config.Server.Timeouts.Write,
		IdleTimeout:     config.Server.Timeouts.Idle,
		Logger:          LoggerPrintf(&ctx.Logger, zerolog.WarnLevel),
	}

	address := net.JoinHostPort(config.Server.Host, strconv.Itoa(config.Server.Port))

	if listener, err = net.Listen("tcp", address); err != nil {
		return nil, nil, fmt.Errorf("unable to initialize tcp listener: %w", err)
	}

	return server, listener, nil
}
