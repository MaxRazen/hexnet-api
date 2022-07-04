package common

import (
	locale_en "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"log"
	"math/rand"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

type ValidationError struct {
	Errors map[string]string `json:"errors"`
}

func GetValidator() *validator.Validate {
	if validate == nil {
		validate = validator.New()

		en := locale_en.New()
		uni = ut.New(en, en)

		trans, _ := uni.GetTranslator("en")
		err := en_translations.RegisterDefaultTranslations(validate, trans)

		if err != nil {
			log.Fatal("i18n Translator could not be initialized")
		}
	}
	return validate
}

func IsValidationError(err error) bool {
	_, ok := err.(validator.ValidationErrors)
	return ok
}

func NewValidatorError(err error) ValidationError {
	res := ValidationError{}

	res.Errors = make(map[string]string)
	errs := err.(validator.ValidationErrors)

	trans, _ := uni.GetTranslator("en")
	translated := errs.Translate(trans)

	for k, v := range translated {
		res.Errors[k] = v
	}

	return res
}

func GenerateRandomString(strLen uint8, allowedCharSet string) string {
	var charSet []rune

	if len(allowedCharSet) > 0 {
		charSet = []rune(allowedCharSet)
	} else {
		charSet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	}

	b := make([]rune, strLen)
	for i := range b {
		b[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(b)
}
