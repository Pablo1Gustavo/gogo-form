package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"gogo-form/models"
	"gogo-form/repository"
)

type FormHandler struct {
	formRepo *repository.FormRepository
}

func NewFormHandler() FormHandler  {
	return FormHandler{repository.NewFormRepository()}
}

func (h *FormHandler) Create(ctx *fiber.Ctx) error {
	form := new(models.Form)
	
	if err := ctx.BodyParser(form); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	form.ID = primitive.NewObjectID()

	_, err := h.formRepo.Create(ctx.Context(), *form)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Could not create the form",
		})
	}

	return ctx.Status(202).JSON(form)
}

func (h *FormHandler) GetAll(ctx *fiber.Ctx) error {
	forms, err := h.formRepo.GetAll(ctx.Context())

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Unexpected error during get forms",
		})
	}

	return ctx.Status(200).JSON(forms)
}

func (h *FormHandler) GetOne(ctx *fiber.Ctx) error {
	id, _ := primitive.ObjectIDFromHex(ctx.Params("id"))

	form, err := h.formRepo.GetOne(ctx.Context(), id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Form not found",
		})
	}

	return ctx.Status(200).JSON(form)
}