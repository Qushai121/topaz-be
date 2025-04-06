package middlewares

import (
	"net/http"
	"strings"

	"github.com/Qushai121/topaz-be/entities"
	"github.com/Qushai121/topaz-be/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthorizationTokenMiddleware(c *fiber.Ctx) error {

	var reqHeaderData = entities.ReqHeader{}

	if err := c.ReqHeaderParser(&reqHeaderData); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	token := strings.Split(reqHeaderData.Authorization, " ")[1]

	encodeToken, errEncodeToken := utils.EncodeToken[entities.UserToken](token, utils.AccessTokenKey)

	if errEncodeToken != nil {
		return errEncodeToken.SendErrorResponse(c)
	}

	c.Locals(entities.AuthorizationToken, encodeToken)

	return c.Next()
}
