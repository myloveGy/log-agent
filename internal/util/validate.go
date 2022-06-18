package util

import (
	"strings"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

type ValidationError struct {
	Field     string `json:"field"`
	Tag       string `json:"tag"`
	Param     string `json:"value"`
	Translate string `json:"translate"`
}

type ErrorResponse struct {
	Errors []*ValidationError `json:"items"`
}

func (e *ErrorResponse) Error() string {
	arr := []string{}
	for _, item := range e.Errors {
		arr = append(arr, item.Translate)
	}

	return strings.Join(arr, ",")
}

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	local := zh.New()
	uni = ut.New(local, local)
	trans, _ = uni.GetTranslator("zh")
	validate = validator.New()
	_ = zhTranslations.RegisterDefaultTranslations(validate, trans)
}

func Validate(data interface{}) *ErrorResponse {
	if err := validate.Struct(data); err != nil {
		errorResponse := &ErrorResponse{
			Errors: make([]*ValidationError, 0),
		}

		for _, err := range err.(validator.ValidationErrors) {
			errorResponse.Errors = append(errorResponse.Errors, &ValidationError{
				Field:     err.StructNamespace(),
				Tag:       err.Tag(),
				Param:     err.Param(),
				Translate: err.Translate(trans),
			})
		}

		return errorResponse
	}

	return nil
}
