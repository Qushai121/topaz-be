package routes

import (
	"github.com/Qushai121/topaz-be/http/controllers"
)

func (route *route) DocumentRoute(documentController controllers.IDocumentController) {
	route.app.Get("/document/list", documentController.GetDocumentList)
}
