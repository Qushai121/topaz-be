package dto

import "github.com/gofiber/fiber/v2"

type SuccessDto[T any] struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    T     `json:"data,omitempty"`
}

func NewSuccessDto[T any](message string, status int, data T) *SuccessDto[T] {
	return &SuccessDto[T]{
		Message: message,
		Status:  status,
		Data:    data,
	}
}

func (e *SuccessDto[T]) SendSuccessResponse(ctx *fiber.Ctx) error {
	return ctx.JSON(e)
}