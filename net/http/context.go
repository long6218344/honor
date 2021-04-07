package http

import (
	"math"
	"net/http"
	"sync"
)

const abortIndex int8 = math.MaxInt8 / 2

type Context struct {
	writermem responseWriter
	Request   *http.Request
	Writer    ResponseWriter

	Params   Params
	handlers HandlersChain
	index    int8
	fullPath string

	engine *Engine
	params *Params

	// This mutex protect Keys map
	mu sync.RWMutex

	// Keys is a key/value pair exclusively for the context of eache request.
	Keys map[string]interface{}
}

func (c *Context) reset() {
}
