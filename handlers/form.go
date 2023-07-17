package handlers

import (
	"github.com/gin-gonic/gin"

	"gogo-form/domain"
	"gogo-form/helpers"
	"gogo-form/repository"
)

type FormHandler struct {
	formRepo domain.FormRepository
}

func NewFormHandler() FormHandler {
	return FormHandler{repository.NewFormRepository()}
}

func (h *FormHandler) Create(ctx *gin.Context) {
	form := new(domain.Form)

	if err := ctx.ShouldBindJSON(form); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Cannot parse JSON",
		})
		return
	}

	err := helpers.ValidateStruct(form)
	if helpers.RespondToError(ctx, err) {
		return
	}

	_, err = h.formRepo.Create(ctx.Request.Context(), *form)
	if helpers.RespondToError(ctx, err) {
		return
	}

	ctx.JSON(202, form)
}

func (h *FormHandler) GetAll(ctx *gin.Context) {
	forms, err := h.formRepo.GetAll(ctx.Request.Context(), ctx.Query("name"))

	if helpers.RespondToError(ctx, err) {
		return
	}

	ctx.JSON(200, forms)
}

func (h *FormHandler) GetOne(ctx *gin.Context) {
	form, err := h.formRepo.GetOne(ctx.Request.Context(), ctx.Param("id"))

	if helpers.RespondToError(ctx, err) {
		return
	}

	ctx.JSON(200, form)
}

func (h *FormHandler) Update(ctx *gin.Context) {
	form := new(domain.Form)

	if err := ctx.ShouldBindJSON(form); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Cannot parse JSON",
		})
		return
	}

	err := helpers.ValidateStruct(form)
	if helpers.RespondToError(ctx, err) {
		return
	}

	updatedForm, err := h.formRepo.Update(ctx.Request.Context(), *form, ctx.Param("id"))

	if helpers.RespondToError(ctx, err) {
		return
	}

	ctx.JSON(200, updatedForm)
}

func (h *FormHandler) Delete(ctx *gin.Context) {
	err := h.formRepo.Delete(ctx.Request.Context(), ctx.Param("id"))

	if helpers.RespondToError(ctx, err) {
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Form successfully deleted",
	})
}
