package middleware

import (
	http2 "github.com/zhoushuguang/honor/frame/http"
)

func Recover() http2.HandlerFunc {
	return func(ctx *http2.Context) {
		defer func() {
			if err := recover(); err != nil {

			}
		}()
		ctx.Next()
	}
}
