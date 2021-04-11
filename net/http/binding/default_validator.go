package binding

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	return nil
}

func (v *defaultValidator) Engine() interface{} {
	return nil
}
