package validators

import (
	"github.com/go-playground/validator/v10"
)

var Val *validator.Validate

func Validates(){
	Val=validator.New()
}