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
		formRoutes.Post("", formController.Create)
		formRoutes.Get("", formController.GetAll)
		formRoutes.Get(":id", formController.GetOne)
	}
	
	answerController := controllers.NewAnswerController()
	answerRoutes := app.Group("answer")
	{
		answerRoutes.Post(":formId", answerController.Create)
		answerRoutes.Get("", answerController.GetAll)
		answerRoutes.Get(":id", answerController.GetOne)
	}

	app.Listen(":3000")
}
