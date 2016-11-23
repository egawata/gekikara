package controllers

import (
	"github.com/revel/revel"
)

func ConvertErrorI18n(c App, v *revel.Validation) {
	for _, verr := range v.Errors {
		verr.Message = c.Message("error." + verr.Message)
	}
}
