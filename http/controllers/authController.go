package controllers

import (
	"net/http"

	"github.com/Qushai121/topaz-be/dto"
	authdto "github.com/Qushai121/topaz-be/dto/authDto"
	"github.com/Qushai121/topaz-be/services"
	"github.com/Qushai121/topaz-be/utils"
	"github.com/gofiber/fiber/v2"
)

type IAuthController interface {
	SignIn(ctx *fiber.Ctx) error
	SignUp(ctx *fiber.Ctx) error
	PostNewAccessToken(ctx *fiber.Ctx) error
}

type authController struct {
	authService services.IAuthService
}

func NewAuthController(authService services.IAuthService) IAuthController {
	return &authController{
		authService: authService,
	}
}

func (a *authController) SignIn(ctx *fiber.Ctx) error {
	body, err := utils.ValidateRequestBody[authdto.SignInRequestBodyDto](ctx)
	if err != nil {
		return err.SendErrorResponse(ctx)
	}

	return dto.NewSuccessDto("Successfully Sign In", http.StatusOK, body).SendSuccessResponse(ctx)
}

func (a *authController) SignUp(ctx *fiber.Ctx) error {
	body, err := utils.ValidateRequestBody[authdto.SignUpRequestBodyDto](ctx)

	if err != nil {
		return err.SendErrorResponse(ctx)
	}

	resService, errService := a.authService.SignUp(body)

	if errService != nil {
		return errService.SendErrorResponse(ctx)
	}

	return dto.NewSuccessDto("Successfully Sign Up", http.StatusOK, resService).SendSuccessResponse(ctx)
}

func (a *authController) PostNewAccessToken(ctx *fiber.Ctx) error {
	body, err := utils.ValidateRequestBody[authdto.PostNewAccessTokenRequestBodyDto](ctx)

	if err != nil {
		return err.SendErrorResponse(ctx)
	}

	resService, errService := a.authService.PostNewAccessToken(body)

	if errService != nil {
		return errService.SendErrorResponse(ctx)
	}

	return resService.SendSuccessResponse(ctx)
}
