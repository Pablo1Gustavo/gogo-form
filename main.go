package main

import (
	"gogo-form/database"

	"gogo-form/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDB()
	app := fiber.New()

	formController := controllers.NewFormController()

	app.Get("", formController.GetAll)
	app.Post("", formController.Create)

	app.Listen(":3000")
}
