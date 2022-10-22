package server

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"matrix-alertmanager/internal/application"
	"matrix-alertmanager/internal/server/handler"
)

func requestRouter(ctx application.Context) fasthttp.RequestHandler {
	r := router.New()
	handlers := handler.New(&ctx)

	r.POST("/alert", handlers.AlertHandler)

	return r.Handler
}
