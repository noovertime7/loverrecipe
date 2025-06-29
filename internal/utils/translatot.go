package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

func GetTrans(locale string) (ut.Translator, error) {
	var trans ut.Translator
	var err error

	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// Register a custom tag name function to get the json tag name
		validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		// Initialize translators
		zhT := zh.New()
		enT := en.New()
		uniT := ut.New(zhT, zhT, enT)

		// Get the translator for the specified locale
		trans, ok = uniT.GetTranslator(locale)
		if !ok {
			return nil, fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// Register the appropriate translations based on the locale
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(validate, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(validate, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(validate, trans)
		}

		if err != nil {
			return nil, fmt.Errorf("registering translations for locale %s failed: %v", locale, err)
		}

		return trans, nil
	}

	return nil, fmt.Errorf("failed to get validator engine")
}
