package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"gogo-form/models"
	"gogo-form/repository"
)

type AnswerController struct {
	answerRepo *repository.AnswerRepository
}

func NewAnswerController() AnswerController {
	answerRepo := repository.NewAnswerRepository()
	return AnswerController{answerRepo}
}

type Answers []interface{}

func (c *AnswerController) Create(ctx *fiber.Ctx) error {
	var answers Answers

	if err := ctx.BodyParser(&answers); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	formId, _ := primitive.ObjectIDFromHex(ctx.Params("formId"))

	_, err := repository.NewFormRepository().GetOne(ctx.Context(), formId)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Form not found",
		})
	}

	answer := models.Answer{
		ID:        primitive.NewObjectID(),
		FormID:    formId,
		AnsweredAt: time.Now(),
		Answers:   answers,
	}

	_, err = c.answerRepo.Create(ctx.Context(), answer)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Could not create answer",
		})
	}

	return ctx.Status(202).JSON(answer)
}

func (c *AnswerController) GetAll(ctx *fiber.Ctx) error {
	forms, err := c.answerRepo.GetAll(ctx.Context())

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Unexpected error during get forms",
		})
	}

	return ctx.Status(200).JSON(forms)
}