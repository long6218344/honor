package http

import "testing"

func TestInsertChild(t *testing.T) {
	node := &node{}
	path := "/user/:name/*addr"
	fullPath := "/user/:name/*addr"

	handlersChain := []HandlerFunc{
		func(ctx *Context) {

		},
		func(ctx *Context) {

		},
	}

	// path: /user/:name/*addr wildcard: :name i: 6 valid: true
	// path: /*addr wildcard: *addr i: 1 valid: true

	node.insertChild(path, fullPath, handlersChain)
}
