package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"gogo-form/models"
	"gogo-form/repository"
)

type AnswerHandler struct {
	answerRepo *repository.AnswerRepository
}

func NewAnswerHandler() AnswerHandler {
	return AnswerHandler{repository.NewAnswerRepository()}
}

type Answers []interface{}

func (h *AnswerHandler) Create(ctx *fiber.Ctx) error {
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

	_, err = h.answerRepo.Create(ctx.Context(), answer)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Could not create answer",
		})
	}

	return ctx.Status(202).JSON(answer)
}

func (h *AnswerHandler) GetAll(ctx *fiber.Ctx) error {
	forms, err := h.answerRepo.GetAll(ctx.Context())

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Unexpected error during get forms",
		})
	}

	return ctx.Status(200).JSON(forms)
}

func (h *AnswerHandler) GetOne(ctx *fiber.Ctx) error {
	id, _ := primitive.ObjectIDFromHex(ctx.Params("id"))

	answer, err := h.answerRepo.GetOne(ctx.Context(), id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Form not found",
		})
	}

	return ctx.Status(200).JSON(answer)
}
