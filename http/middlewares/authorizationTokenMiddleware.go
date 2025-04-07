package middlewares

import (
	"net/http"
	"strings"

	"github.com/Qushai121/topaz-be/dto"
	"github.com/Qushai121/topaz-be/entities"
	"github.com/Qushai121/topaz-be/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthorizationTokenMiddleware(c *fiber.Ctx) error {

	var reqHeaderData = entities.ReqHeader{}

	if err := c.ReqHeaderParser(&reqHeaderData); err != nil {
		return dto.NewErrorDto[any](err.Error(), http.StatusUnauthorized, nil).SendErrorResponse(c)
	}

	if reqHeaderData.Authorization == nil {
		return dto.NewErrorDto[any]("Authorization Header not found", http.StatusUnauthorized, nil).SendErrorResponse(c)
	}

	token := strings.Split(*reqHeaderData.Authorization, " ")[1]

	encodeToken, errEncodeToken := utils.EncodeToken[entities.UserTokenPayload](token, utils.AccessTokenKey)

	if errEncodeToken != nil {
		return errEncodeToken.SendErrorResponse(c)
	}

	c.Locals(entities.AuthorizationToken, encodeToken.Data)

	return c.Next()
}
