package main

import (
	"log"

	"github.com/zhoushuguang/honor/net/http"
)

func main() {
	engine := http.New()
	engine.GET("/user", func(ctx *http.Context) {
		ctx.Writer.Write([]byte("hello honor"))
	})

	log.Fatal(engine.Run(":9000"))
}
