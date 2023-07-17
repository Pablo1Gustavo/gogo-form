package main

import (
	"gogo-form/database"
	"gogo-form/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	app := gin.Default()

	formHandler := handlers.NewFormHandler()
	formRoutes := app.Group("/form")
	{
		formRoutes.POST("", formHandler.Create)
		formRoutes.GET("", formHandler.GetAll)
		formRoutes.GET(":id", formHandler.GetOne)
		formRoutes.PUT(":id", formHandler.Update)
		formRoutes.DELETE(":id", formHandler.Delete)
	}

	answerHandler := handlers.NewAnswerHandler()
	answerRoutes := app.Group("/answer")
	{
		answerRoutes.POST(":formId", answerHandler.Create)
		answerRoutes.GET("", answerHandler.GetAll)
		answerRoutes.GET(":id", answerHandler.GetOne)
	}

	app.Run(":3000")
}
