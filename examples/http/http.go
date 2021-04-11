package main

import (
	"github.com/zhoushuguang/honor/net/http"
)

func main() {
	engine := &http.Engine{}
	engine.GET("/user/info", func(ctx *http.Context) {

	})
	engine.Run()
}
