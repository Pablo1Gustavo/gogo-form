package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"gogo-form/models"
	"gogo-form/repository"
)

type formController struct {
	formRepo *repository.FormRepository
}

func NewFormController() formController  {
	formRepo := repository.NewFormRepository()
	return formController{formRepo}
}

func (c *formController) Create(ctx *fiber.Ctx) error {
	form := new(models.Form)
	
	if err := ctx.BodyParser(form); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	form.ID = primitive.NewObjectID()

	_, err := c.formRepo.Create(ctx.Context(), *form)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Could not create the form",
		})
	}

	return ctx.Status(202).JSON(form)
}

func (c *formController) GetAll(ctx *fiber.Ctx) error {
	forms, err := c.formRepo.GetAll(ctx.Context())

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Unexpected error during get forms",
		})
	}

	return ctx.Status(200).JSON(forms)
}

func (c *formController) GetOne(ctx *fiber.Ctx) error {
	form, err := c.formRepo.GetOne(ctx.Context(), ctx.Params("id"))
	
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Form not found",
		})
	}

	return ctx.Status(200).JSON(form)
}