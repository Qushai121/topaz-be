package utils

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	id_translations "github.com/go-playground/validator/v10/translations/id"

	"github.com/Qushai121/topaz-be/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Language string

const (
	EN = "en"
	ID = "id"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func InitValidate(lang Language) {

	validate = validator.New(validator.WithRequiredStructEnabled())
	en := en.New()
	id := id.New()
	uni = ut.New(en, en, id)

	switch lang {
	case EN:
		{
			trans, _ = uni.GetTranslator("en")
			en_translations.RegisterDefaultTranslations(validate, trans)
		}
	case ID:
		{
			trans, _ = uni.GetTranslator("id")
			id_translations.RegisterDefaultTranslations(validate, trans)
		}
	}
}

func ValidateRequestBody[T any](ctx *fiber.Ctx) (*T, *dto.ErrorDto[any]) {
	var body T
	if err := ctx.BodyParser(&body); err != nil {
		log.Println(err.Error())
		return nil, dto.BadRequestError()
	}

	err := validate.Struct(&body)

	errFields := make(map[string][]string)

	if err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			field := e.Field()
			field = strings.ToLower(string(field[0])) + field[1:]

			errFields[field] = append(errFields[field], e.Translate(trans))
		}
	}

	if len(errFields) > 0 {
		return nil, dto.NewErrorDto("Request body doesn`t meet the need", http.StatusUnprocessableEntity, any(&errFields))
	}

	return &body, nil
}

func Coba() {

}
