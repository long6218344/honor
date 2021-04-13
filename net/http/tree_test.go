package http

import (
	"fmt"
	"testing"
)

func TestInsertChild(t *testing.T) {
	node := &node{}
	path := "/us/:city/*addr"
	fullPath := "/us/:city/*addr"

	handlersChain := []HandlerFunc{
		func(ctx *Context) {

		},
		func(ctx *Context) {

		},
	}

	// path: /user/:name/*addr wildcard: :name i: 6 valid: true
	// path: /*addr wildcard: *addr i: 1 valid: true

	node.insertChild(path, fullPath, handlersChain)

	fmt.Printf("%+v\n", node)
	fmt.Printf("%+v\n", node.children[0])
	fmt.Printf("%+v\n", node.children[0].children[0])
	fmt.Printf("%+v\n", node.children[0].children[0].children[0])
	fmt.Printf("%+v\n", node.children[0].children[0].children[0].children[0])
}
