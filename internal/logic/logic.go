package logic

import (
	"matrix-alertmanager/internal/application"
)

type logic struct {
	ctx *application.Context
}

func New(ctx *application.Context) application.Logic {
	return &logic{ctx: ctx}
}
