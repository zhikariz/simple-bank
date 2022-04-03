package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/zhikariz/simple-bank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrenncy(currency)
	}
	return false
}