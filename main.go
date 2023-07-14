package main

import (
	"gogo-form/database"
	"gogo-form/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDB()
	app := fiber.New()

	formHandler := handlers.NewFormHandler()
	formRoutes := app.Group("form")
	{
		formRoutes.Post("", formHandler.Create)
		formRoutes.Get("", formHandler.GetAll)
		formRoutes.Get(":id", formHandler.GetOne)
	}
	
	answerHandler := handlers.NewAnswerHandler()
	answerRoutes := app.Group("answer")
	{
		answerRoutes.Post(":formId", answerHandler.Create)
		answerRoutes.Get("", answerHandler.GetAll)
		answerRoutes.Get(":id", answerHandler.GetOne)
	}

	app.Listen(":3000")
}
