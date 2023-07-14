package handlers

import (
	"github.com/gofiber/fiber/v2"

	"gogo-form/domain"
	"gogo-form/repository"
)

type FormHandler struct {
	formRepo *repository.FormRepository
}

func NewFormHandler() FormHandler  {
	return FormHandler{repository.NewFormRepository()}
}

func (h *FormHandler) Create(ctx *fiber.Ctx) error {
	form := new(domain.Form)
	
	if err := ctx.BodyParser(form); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	_, err := h.formRepo.Create(ctx.Context(), *form)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Could not create the form",
		})
	}

	return ctx.Status(202).JSON(form)
}

func (h *FormHandler) GetAll(ctx *fiber.Ctx) error {
	forms, err := h.formRepo.GetAll(ctx.Context(), ctx.Query("name"))

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Unexpected error during get forms",
		})
	}

	return ctx.Status(200).JSON(forms)
}

func (h *FormHandler) GetOne(ctx *fiber.Ctx) error {
	form, err := h.formRepo.GetOne(ctx.Context(), ctx.Params("id"))
	
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Form not found",
		})
	}

	return ctx.Status(200).JSON(form)
}