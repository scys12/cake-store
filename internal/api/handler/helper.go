package handler

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func ValidateRequest(validate *validator.Validate, err error) (errs []error) {
	if err == nil {
		return nil
	}

	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)

	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}

func GetValidatedSortBy(sortStr string) string {
	if sortStr != "title" && sortStr != "rating" && sortStr != "created_at" {
		return ""
	}
	return sortStr
}
