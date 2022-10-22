package handler

import (
	"matrix-alertmanager/internal/application"
)

type Handlers struct {
	ctx *application.Context
}

func New(ctx *application.Context) Handlers {
	return Handlers{ctx: ctx}
}
