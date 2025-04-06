package routes

import (
	"github.com/Qushai121/topaz-be/http/controllers"
)

func (route *route) AuthRoute(authController controllers.IAuthController) {
	route.app.Post("/sign-in", authController.SignIn)
	route.app.Post("/sign-up", authController.SignUp)
	route.app.Post("/refresh", authController.PostNewAccessToken)
}
