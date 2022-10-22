package server

import "github.com/valyala/fasthttp"

func handleError() func(ctx *fasthttp.RequestCtx, err error) {
	return func(ctx *fasthttp.RequestCtx, err error) {

	}
}
