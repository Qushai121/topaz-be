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
	"github.com/Qushai121/topaz-be/entities"
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

// Validating and getting return Request body using go validator
func ValidateRequestBody[T any](ctx *fiber.Ctx) (*T, *dto.ErrorDto[any]) {
	var body T
	if err := ctx.BodyParser(&body); err != nil {
		log.Println(err.Error())
		return nil, dto.BadRequestError()
	}

	errFields := ValidateStruct(body)

	if len(errFields) > 0 {
		return nil, dto.NewErrorDto("Request body doesn`t meet the need", http.StatusUnprocessableEntity, any(&errFields))
	}

	return &body, nil
}

// Validating and getting return query params using go validator
func ValidateQueryParams[T any](ctx *fiber.Ctx) (*T, *dto.ErrorDto[any]) {
	var queryParams T
	if err := ctx.QueryParser(&queryParams); err != nil {
		log.Println(err.Error())
		return nil, dto.BadRequestError()
	}

	errFields := ValidateStruct(queryParams)

	if len(errFields) > 0 {
		return nil, dto.NewErrorDto("Request query params doesn`t meet the need", http.StatusUnprocessableEntity, any(&errFields))
	}

	return &queryParams, nil
}

// Validating and getting return params / path url using go validator
func ValidateParams[T any](ctx *fiber.Ctx) (*T, *dto.ErrorDto[any]) {
	var params T
	if err := ctx.ParamsParser(&params); err != nil {
		log.Println(err.Error())
		return nil, dto.BadRequestError()
	}

	errFields := ValidateStruct(params)

	if len(errFields) > 0 {
		return nil, dto.NewErrorDto("params doesn`t meet the need", http.StatusUnprocessableEntity, any(&errFields))
	}

	return &params, nil
}

// Validate and getting request body multipart
// multipart ussualy have file upload so it mandtory to specify file fields so parse multipart can assign it to body request
func ValidateRequestBodyMultipart[T any](ctx *fiber.Ctx, fileFields *[]entities.FileField) (*T, *dto.ErrorDto[any]) {
	var body T
	if err := ParseMultipartRequest(ctx, &body, fileFields); err != nil {
		log.Println(err.Error())
		return nil, dto.BadRequestError()
	}

	errFields := ValidateStruct(body)

	if len(errFields) > 0 {
		return nil, dto.NewErrorDto("Request body doesn`t meet the need", http.StatusUnprocessableEntity, any(&errFields))
	}

	return &body, nil
}

func ValidateStruct[T any](data T) map[string][]string {

	err := validate.Struct(&data)

	errFields := make(map[string][]string)

	if err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			field := e.Field()
			field = strings.ToLower(string(field[0])) + field[1:]
			errFields[field] = append(errFields[field], e.Translate(trans))
		}
	}

	return errFields
}
