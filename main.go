package main

import (
	"log"

	"github.com/Qushai121/topaz-be/configs"
	"github.com/Qushai121/topaz-be/http/controllers"
	"github.com/Qushai121/topaz-be/utils"

	// "github.com/Qushai121/topaz-be/models"
	"github.com/Qushai121/topaz-be/routes"
	"github.com/Qushai121/topaz-be/services"
	fiber "github.com/gofiber/fiber/v2"
	recover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	initAll()
	app := fiber.New()
	app.Use(recover.New())

	routes := routes.NewRoute(app)

	// Document Route Controller Inject Document Service
	routes.DocumentRoute(controllers.NewDocumentController(services.NewDocumentService(configs.DBTOPAZ)))
	routes.AuthRoute(controllers.NewAuthController(services.NewAuthService(configs.DBTOPAZ)))
	app.Listen(":3000")
}

func initAll() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.InitValidate(utils.ID)
	configs.InitDBTopaz()

	// configs.InitDBTopaz.AutoMigrate(
	// 	&models.User{},
	// 	&models.Notification{},
	// 	&models.News{},
	// 	&models.BannerInfo{},
	// 	&models.CategoryDocument{},
	// 	&models.ContentDocument{},
	// 	&models.Document{},
	// 	&models.DocumentFileStorage{},
	// 	&models.FileStorage{},
	// 	&models.Role{},
	// 	&models.RoleUser{},
	// )
}
