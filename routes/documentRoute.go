package routes

import (
	"github.com/Qushai121/topaz-be/http/controllers"
	"github.com/Qushai121/topaz-be/http/middlewares"
)

func (route *route) DocumentRoute(documentController controllers.IDocumentController) {

	route.app.Get("/document/list", documentController.GetDocumentList)
	authorized := route.app.Group("/document", middlewares.AuthorizationTokenMiddleware)
	authorized.Post("", documentController.CreateDocument)
	authorized.Put("/:id", documentController.UpdateDocument)
	authorized.Delete("/:id", documentController.DeleteDocument)
}
