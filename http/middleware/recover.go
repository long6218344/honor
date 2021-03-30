package middleware

import "github.com/zhoushuguang/honor/http"

func Recover() http.HandlerFunc {
	return func(ctx *http.Context) {
		defer func() {
			if err := recover(); err != nil {

			}
		}()
		ctx.Next()
	}
}
