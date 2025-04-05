package routes

import (
	"github.com/Qushai121/topaz-be/http/controllers"
	"github.com/gofiber/fiber/v2"
)

type IRoute interface {
	DocumentRoute(documentController controllers.IDocumentController)
	AuthRoute(authController controllers.IAuthController)
}

type route struct {
	app *fiber.App
}

func NewRoute(app *fiber.App) IRoute {
	return &route{app: app}
}
