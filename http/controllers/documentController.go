package controllers

import (
	"fmt"

	documentdto "github.com/Qushai121/topaz-be/dto/documentDto"
	"github.com/Qushai121/topaz-be/services"
	"github.com/Qushai121/topaz-be/utils"
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
	query, err := utils.ValidateQueryParams[documentdto.GetDocumentListQueryParamsDto](ctx)

	fmt.Println(query.Page)
	if err != nil {
		return err.SendErrorResponse(ctx)
	}

	resService, errService := d.documentServices.GetDocumentList(query)

	if errService != nil {
		return errService.SendErrorResponse(ctx)
	}

	return resService.SendSuccessResponse(ctx)
}
