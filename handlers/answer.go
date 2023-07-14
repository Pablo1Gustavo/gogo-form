package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"gogo-form/domain"
	"gogo-form/repository"
)

type AnswerHandler struct {
	answerRepo *repository.AnswerRepository
}

func NewAnswerHandler() AnswerHandler {
	return AnswerHandler{repository.NewAnswerRepository()}
}

func (h *AnswerHandler) Create(ctx *fiber.Ctx) error {
	var answers []interface{}

	if err := ctx.BodyParser(&answers); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	formID := ctx.Params("formId")

	form, err := repository.NewFormRepository().GetOne(ctx.Context(), formID)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Form not found",
		})
	}

	formAnswer := domain.Answer{
		FormID:     formID,
		AnsweredAt: time.Now(),
		Answers:    answers,
	}

	if !formAnswer.CompatibleWithForm(form) {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "The response is not compatible with the form.",
		})
	}

	formAnswer, err = h.answerRepo.Create(ctx.Context(), formAnswer)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Could not create answer",
		})
	}

	return ctx.Status(202).JSON(formAnswer)
}

func (h *AnswerHandler) GetAll(ctx *fiber.Ctx) error {
	answers, err := h.answerRepo.GetAll(ctx.Context())

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Unexpected error during get forms",
		})
	}

	return ctx.Status(200).JSON(answers)
}

func (h *AnswerHandler) GetOne(ctx *fiber.Ctx) error {
	formAnswer, err := h.answerRepo.GetOne(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Form not found",
		})
	}

	return ctx.Status(200).JSON(formAnswer)
}
