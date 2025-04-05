package controllers

import (
	"github.com/Qushai121/topaz-be/services"
	"github.com/gofiber/fiber/v2"
)

type IDocumentController interface {
	GetDocumentList(ctx *fiber.Ctx) error
}

type documentController struct {
	documentServices services.IDocumentService
}

func NewDocumentController(documentServices services.IDocumentService) IDocumentController {
	return &documentController{
		documentServices: documentServices,
	}
}

func (d *documentController) GetDocumentList(ctx *fiber.Ctx) error {
	return nil
}
