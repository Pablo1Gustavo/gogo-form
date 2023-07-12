package main

import (
	"gogo-form/controllers"
	"gogo-form/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDB()
	app := fiber.New()

	formController := controllers.NewFormController()

	formRoutes := app.Group("form")
	{
		formRoutes.Get("", formController.GetAll)
		formRoutes.Post("", formController.Create)
		formRoutes.Get(":id", formController.GetOne)
	}

	app.Listen(":3000")
}
