package dto

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ErrorDto[T any] struct {
	Message string `json:"message"`
	Status  int `json:"status"`
 	Data    T `json:"data"`
}

func (e *ErrorDto[T]) SendErrorResponse(ctx *fiber.Ctx) error {
	return ctx.JSON(e)
}

func NewErrorDto[T any](message string, status int, data T) *ErrorDto[T] {
	return &ErrorDto[T]{
		Message: message,
		Status:  status,
		Data:    data,
	}
}

func InternalServerError() *ErrorDto[any] {
	return NewErrorDto[any](
		"Internal Server Error",
		http.StatusInternalServerError,
		nil,
	)
}

func UnathorizedActionError() *ErrorDto[any] {
	return NewErrorDto[any](
		"Unathorized Action Error",
		http.StatusUnauthorized,
		nil,
	)
}
