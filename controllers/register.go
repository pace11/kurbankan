package controllers

import (
	"kurbankan/models"
	"kurbankan/repository"
	"kurbankan/utils"

	"github.com/gin-gonic/gin"
)

type RegisterController struct {
	Repo repository.RegisterRepository
}

func NewRegisterController(repo repository.RegisterRepository) *RegisterController {
	return &RegisterController{Repo: repo}
}

func (ctl *RegisterController) RegisterParticipant(ctx *gin.Context) {
	var payload models.UserCreatePayload

	if utils.BindAndValidate(ctx, &payload) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.SaveParticipant(&payload)
	utils.HttpResponse(ctx, data, code, entity, ctx.Request.Method, errors)
}

func (ctl *RegisterController) RegisterMosque(ctx *gin.Context) {
	var payload models.UserCreatePayload

	if utils.BindAndValidate(ctx, &payload) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.SaveMosque(&payload)
	utils.HttpResponse(ctx, data, code, entity, ctx.Request.Method, errors)
}
