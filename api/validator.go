package api

import (
	"github.com/HectorSauR/simplebank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {

	if currency, ok := fl.Field().Interface().(string); ok {
		//check if currrency is supported
		return util.IsSupportedCurrency(currency)
	}

	return false
}
