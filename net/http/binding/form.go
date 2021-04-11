package binding

import "net/http"

type formBinding struct{}

func (formBinding) Name() string {
	return "form"
}

func (formBinding) Bind(req *http.Request, obj interface{}) error {
	return nil
}
