package handlers

import (
	"time"

	"github.com/gin-gonic/gin"

	"gogo-form/domain"
	"gogo-form/helpers"
	"gogo-form/repository"
)

type AnswerHandler struct {
	answerRepo domain.AnswerRepository
}

func NewAnswerHandler() AnswerHandler {
	return AnswerHandler{repository.NewAnswerRepository()}
}

func (h *AnswerHandler) Create(ctx *gin.Context) {
	var answers []interface{}

	if err := ctx.ShouldBindJSON(&answers); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Cannot parse JSON",
		})
		return
	}

	formID := ctx.Param("formId")

	form, err := repository.NewFormRepository().GetOne(ctx.Request.Context(), formID)
	if helpers.RespondToError(ctx, err) {
		return
	}

	formAnswer := domain.Answer{
		FormID:     formID,
		AnsweredAt: time.Now(),
		Answers:    answers,
	}

	if !formAnswer.CompatibleWithForm(form) {
		ctx.JSON(422, gin.H{
			"message": "The response is not compatible with the form.",
		})
		return
	}

	formAnswer, err = h.answerRepo.Create(ctx.Request.Context(), formAnswer)
	if helpers.RespondToError(ctx, err) {
		return
	}

	ctx.JSON(202, formAnswer)
}

func (h *AnswerHandler) GetAll(ctx *gin.Context) {
	answers, err := h.answerRepo.GetAll(ctx.Request.Context())

	if helpers.RespondToError(ctx, err) {
		return
	}

	ctx.JSON(200, answers)
}

func (h *AnswerHandler) GetOne(ctx *gin.Context) {
	formAnswer, err := h.answerRepo.GetOne(ctx.Request.Context(), ctx.Param("id"))

	if helpers.RespondToError(ctx, err) {
		return
	}

	ctx.JSON(200, formAnswer)
}

func (h *AnswerHandler) Delete(ctx *gin.Context) {
	err := h.answerRepo.Delete(ctx.Request.Context(), ctx.Param("id"))

	if helpers.RespondToError(ctx, err) {
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Answer successfully deleted",
	})
}
