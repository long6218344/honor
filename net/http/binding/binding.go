package binding

import (
	"net/http"
)

// Content-Type MIME of the most common data formats.
const (
	MIMEJSON              = "application/json"
	MIMEHTML              = "text/html"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
)

type Binding interface {
	Name() string
	Bind(*http.Request, interface{}) error
}

type StructValidator interface {
	ValidateStruct(interface{}) error
	Engine() interface{}
}

var Validator StructValidator = &defaultValidator{}

var (
	JOSN = jsonBinding{}
	Form = formBinding{}
)

// Default returns the appropriate Binding instance based on the HTTP method
// and the content type.
func Default(method, contentType string) Binding {
	if method == http.MethodGet {
		return Form
	}

	switch contentType {
	case MIMEJSON:
		return JOSN
	default:
		return Form
	}
}

func validate(obj interface{}) error {
	if Validator == nil {
		return nil
	}
	return Validator.ValidateStruct(obj)
}
