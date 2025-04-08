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

	resService, errService := d.documentServices.GetDocumentList(queryParams)
	if errService != nil {
		return errService.SendErrorResponse(ctx)
	}
	return resService.SendSuccessResponse(ctx)
}

func (d *documentController) CreateDocument(ctx *fiber.Ctx) error {
	reqBody, errReqBody := utils.ValidateRequestBodyMultipart[documentdto.DocumentRequestBodyDto](
		ctx,
		&[]entities.FileField{
			{
				Field:      "image",
				IsMultiple: false,
			},
		},
	)
	if errReqBody != nil {
		return errReqBody.SendErrorResponse(ctx)
	}

	localDataUser, errLocalDataUser := utils.GetAuthorizationTokenFromLocals(ctx)
	if errLocalDataUser != nil {
		return errLocalDataUser.SendErrorResponse(ctx)
	}

	resService, errService := d.documentServices.CreateDocument(ctx, reqBody, localDataUser)
	if errService != nil {
		return errService.SendErrorResponse(ctx)
	}
	return resService.SendSuccessResponse(ctx)
}

func (d *documentController) DeleteDocument(ctx *fiber.Ctx) error {
	params, err := utils.ValidateParams[dto.BaseDetailParamsDto](ctx)
	if err != nil {
		return err.SendErrorResponse(ctx)
	}

	localDataUser, errLocalDataUser := utils.GetAuthorizationTokenFromLocals(ctx)
	if errLocalDataUser != nil {
		return errLocalDataUser.SendErrorResponse(ctx)
	}

	resService, errService := d.documentServices.DeleteDocument(params.ID, localDataUser.UserId)
	if errService != nil {
		return errService.SendErrorResponse(ctx)
	}
	return resService.SendSuccessResponse(ctx)
}

func (d *documentController) UpdateDocument(ctx *fiber.Ctx) error {
	// params, err := utils.ValidateParams[dto.BaseDetailParamsDto](ctx)
	// if err != nil {
	// 	return err.SendErrorResponse(ctx)
	// }

	// reqBody, errReqBody := utils.ValidateRequestBody[documentdto.DocumentRequestBodyDto](ctx)
	// if errReqBody != nil {
	// 	return errReqBody.SendErrorResponse(ctx)
	// }

	// localDataUser, errLocalDataUser := utils.GetAuthorizationTokenFromLocals(ctx)
	// if errLocalDataUser != nil {
	// 	return errLocalDataUser.SendErrorResponse(ctx)
	// }

	// reqBody.UserId = localDataUser.UserId

	// resService, errService := d.documentServices.UpdateDocument(params.ID, localDataUser.UserId, reqBody)
	// if errService != nil {
	// 	return errService.SendErrorResponse(ctx)
	// }
	// return resService.SendSuccessResponse(ctx)
	return nil
}
