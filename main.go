package main

import (
	"gogo-form/database"
	"gogo-form/handlers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.InitDB()
	app := gin.Default()

	formHandler := handlers.NewFormHandler()
	answerHandler := handlers.NewAnswerHandler()

	formRoutes := app.Group("/form")
	{
		formRoutes.POST("", formHandler.Create)
		formRoutes.GET("", formHandler.GetAll)
		formRoutes.GET(":id", formHandler.GetOne)
		formRoutes.PUT(":id", formHandler.Update)
		formRoutes.DELETE(":id", formHandler.Delete)

		formRoutes.GET(":id/answer", answerHandler.GetAll)
		formRoutes.POST(":id/answer", answerHandler.Create)
	}
	answerRoutes := app.Group("/answer")
	{
		answerRoutes.GET("", answerHandler.GetAll)
		answerRoutes.GET(":id", answerHandler.GetOne)
		answerRoutes.DELETE(":id", answerHandler.Delete)
	}

	app.Run(":3000")
}
