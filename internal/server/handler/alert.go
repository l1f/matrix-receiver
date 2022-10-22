package handler

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"matrix-alertmanager/internal/alertmanager"
)

func (h Handlers) AlertHandler(ctx *fasthttp.RequestCtx) {
	var webhook alertmanager.Webhook

	err := json.Unmarshal(ctx.Request.Body(), &webhook)
	if err != nil {
		h.ctx.Logger.Error().Err(err).Send()
		ctx.SetBody([]byte("ERR"))
		return
	}

	err = h.ctx.Logic.ScheduleMessage(webhook)
	if err != nil {
		h.ctx.Logger.Error().Err(err).Send()
		ctx.SetBody([]byte("ERR"))
		return
	}

	ctx.SetBody([]byte("OK"))
}
