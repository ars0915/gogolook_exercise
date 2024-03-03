package router

import (
	"github.com/ars0915/gogolook-exercise/util/cGin"
)

func resourceCheck(rH HttpHandler) cGin.HandlerFunc {
	return func(ctx *cGin.Context) {

		ctx.Next()
	}
}
