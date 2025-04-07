package controllers

import (
	"github.com/Qushai121/topaz-be/dto"
	documentdto "github.com/Qushai121/topaz-be/dto/documentDto"
	"github.com/Qushai121/topaz-be/entities"
	"github.com/Qushai121/topaz-be/services"
	"github.com/Qushai121/topaz-be/utils"
	"github.com/gofiber/fiber/v2"
)

type IDocumentController interface {
	GetDocumentList(ctx *fiber.Ctx) error
	CreateDocument(ctx *fiber.Ctx) error
	UpdateDocument(ctx *fiber.Ctx) error
	DeleteDocument(ctx *fiber.Ctx) error
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
	queryParams, err := utils.ValidateQueryParams[documentdto.GetDocumentListQueryParamsDto](ctx)

	if err != nil {
		return err.SendErrorResponse(ctx)
	}

	resService, err := d.documentServices.GetDocumentList(queryParams)

	if err != nil {
		return err.SendErrorResponse(ctx)
	}

	return resService.SendSuccessResponse(ctx)
}

func (d *documentController) CreateDocument(ctx *fiber.Ctx) error {
	reqBody, err := utils.ValidateRequestBody[documentdto.DocumentRequestBodyDto](ctx)

	localDataUser, ok := ctx.Locals(entities.AuthorizationToken).(entities.UserTokenPayload)

	if !ok {
		return dto.BadRequestError().SendErrorResponse(ctx)
	}

	reqBody.UserId = localDataUser.UserId

	if err != nil {
		return err.SendErrorResponse(ctx)
	}

	resService, err := d.documentServices.CreateDocument(reqBody)

	if err != nil {
		return err.SendErrorResponse(ctx)
	}

	return resService.SendSuccessResponse(ctx)
}

func (d *documentController) DeleteDocument(ctx *fiber.Ctx) error {
	params, err := utils.ValidateParams[dto.BaseDetailParamsDto](ctx)

	if err != nil {
		return err.SendErrorResponse(ctx)
	}

	localDataUser, ok := ctx.Locals(entities.AuthorizationToken).(entities.UserTokenPayload)

	if !ok {
		return dto.BadRequestError().SendErrorResponse(ctx)
	}

	resService, err := d.documentServices.DeleteDocument(params.ID, localDataUser.UserId)

	if err != nil {
		return err.SendErrorResponse(ctx)
	}

	return resService.SendSuccessResponse(ctx)
}

func (d *documentController) UpdateDocument(ctx *fiber.Ctx) error {
	params, err := utils.ValidateParams[dto.BaseDetailParamsDto](ctx)

	if err != nil {
		return err.SendErrorResponse(ctx)
	}

	reqBody, errReqBody := utils.ValidateRequestBody[documentdto.DocumentRequestBodyDto](ctx)

	if errReqBody != nil {
		return err.SendErrorResponse(ctx)
	}

	localDataUser, ok := ctx.Locals(entities.AuthorizationToken).(entities.UserTokenPayload)

	if !ok {
		return dto.BadRequestError().SendErrorResponse(ctx)
	}

	resService, err := d.documentServices.UpdateDocument(params.ID, localDataUser.UserId, reqBody)

	if err != nil {
		return err.SendErrorResponse(ctx)
	}

	return resService.SendSuccessResponse(ctx)
}
