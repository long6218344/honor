package http

import (
	"context"
	"math"
	"net/http"
	"net/url"
	"sync"

	"github.com/zhoushuguang/honor/net/http/binding"
)

const abortIndex int8 = math.MaxInt8 / 2

// Context allows us to pass variables between middleware, manage the flow
// Validate the JSON of a request and render a JSON response for example.
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

	// Errors is a list of errors attached to all the handlers/middlewares who used this context.
	Errors errorMsgs

	Accepted []string

	queryCache url.Values
	formCache  url.Values

	ctx context.Context
}

func (c *Context) reset() {
	c.Writer = &c.writermem
	c.Params = c.Params[0:0]
	c.handlers = nil
	c.index = -1

	c.fullPath = ""
	c.Keys = nil
	c.Errors = c.Errors[0:0]
	c.Accepted = nil
	c.queryCache = nil
	c.formCache = nil
	*c.params = (*c.params)[0:0]
}

func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *Context) Param(key string) string {
	return c.Params.ByName(key)
}

func (c *Context) Query(key string) string {
	value, _ := c.GetQuery(key)
	return value
}

func (c *Context) GetQuery(key string) (string, bool) {
	if values, ok := c.GetQueryArray(key); ok {
		return values[0], ok
	}
	return "", false
}

func (c *Context) GetQueryArray(key string) ([]string, bool) {
	c.initQueryCache()
	if values, ok := c.queryCache[key]; ok && len(values) > 0 {
		return values, true
	}
	return []string{}, false
}

func (c *Context) PostForm(key string) string {
	value, _ := c.GetPostForm(key)
	return value
}

func (c *Context) GetPostForm(key string) (string, bool) {
	if values, ok := c.GetPostFormArray(key); ok {
		return values[0], ok
	}
	return "", false
}

func (c *Context) GetPostFormArray(key string) ([]string, bool) {
	c.initFormCache()
	if values := c.formCache[key]; len(values) > 0 {
		return values, true
	}
	return []string{}, false
}

func (c *Context) initQueryCache() {
	if c.queryCache == nil {
		if c.Request != nil {
			c.queryCache = c.Request.URL.Query()
		} else {
			c.queryCache = url.Values{}
		}
	}
}

func (c *Context) initFormCache() {
	if c.formCache == nil {
		c.formCache = make(url.Values)
		req := c.Request
		if err := req.ParseMultipartForm(c.engine.MaxMultipartMemory); err != nil {
			if err != http.ErrNotMultipart {

			}
		}
		c.formCache = req.PostForm
	}
}

func (c *Context) AbortWithError(code int, err error) *Error {
	c.AbortWithStatus(code)
	return c.Error(err)
}

func (c *Context) AbortWithStatus(code int) {
	c.Status(code)
	c.Writer.WriteHeaderNow()
	c.Abort()
}

// Abort prevents pending handlers from being called. Note that this will not stop the current handler.
// Let's say you have an authorization middleare that validates that the current request is authorized.
// If the authorization fails (ex: the password does not match), call Abort to ensure the remaing handlers
// for this request are not called.
func (c *Context) Abort() {
	c.index = abortIndex
}

// Bind checks the Content-Type to select a binding engine automatically,
// Depending the "Content-Type" header different bindings are used:
//		"application/json" --> JSON binding
//		"application/xml"  --> XML binding
// otherwise --> returns an error.
// It parses the request's body as JSON if Content-Type == "application/json" using JSON or XML as a JSON input.
// It decodes the json payload into the struct specified as a pointer.
// It writes a 400 error and sets Content-Type header "text/plain" in the response if input is not valid.
func (c *Context) Bind(obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.MustBindWith(obj, b)
}

// MustBindWith binds the passed struct pointer using the specified binding engine.
// It will abort the request with HTTP 400 if any error occurs.
func (c *Context) MustBindWith(obj interface{}, b binding.Binding) error {
	if err := c.ShouldBindWith(obj, b); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err).SetType(ErrorTypeBind)
		return err
	}
	return nil
}

// ShouldBindWith binds the passed struct pointer using the specified binding engine.
func (c *Context) ShouldBindWith(obj interface{}, b binding.Binding) error {
	return b.Bind(c.Request, obj)
}

func (c *Context) ContentType() string {
	return filterFlags(c.requestHeader("Content-Type"))
}
func (c *Context) requestHeader(key string) string {
	return c.Request.Header.Get(key)
}

// Status sets the HTTP response code.
func (c *Context) Status(code int) {
	c.Writer.WriteHeader(code)
}

// Error attaches an error to the current context. The error is pushed to a list of errors.
// It's a good idea to call Error for each error that occurred during the resolution of a request.
// A middleare can be used to collect all the errors and push them to a database together,
// print a log, or append it in the HTTP response.
// Error will panice if err is nil.
func (c *Context) Error(err error) *Error {
	if err == nil {
		panic("err is nil")
	}

	parsedError, ok := err.(*Error)
	if !ok {
		parsedError = &Error{
			Err:  err,
			Type: ErrorTypePrivate,
		}
	}
	c.Errors = append(c.Errors, parsedError)
	return parsedError
}
