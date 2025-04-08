package utils

import (
	"github.com/Qushai121/topaz-be/dto"
	"github.com/Qushai121/topaz-be/entities"
	"github.com/gofiber/fiber/v2"
)

// Get UserTokenPayload from localc data with named key entities.AuthorizationToken
func GetAuthorizationTokenFromLocals(ctx *fiber.Ctx) (*entities.UserTokenPayload, *dto.ErrorDto[any]) {
	localDataUser, ok := ctx.Locals(entities.AuthorizationToken).(entities.UserTokenPayload)

	if !ok {
		return nil, dto.BadRequestError()
	}

	return &localDataUser, nil

}
