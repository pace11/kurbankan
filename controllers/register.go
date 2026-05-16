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

	data, code, _, errors := ctl.Repo.SaveParticipant(&payload)
	if utils.HandleRepoError(ctx, code, errors) {
		return
	}
	utils.MutationResponse(ctx, code, "Participant registered successfully", data)
}

func (ctl *RegisterController) RegisterMosque(ctx *gin.Context) {
	var payload models.UserCreatePayload

	if utils.BindAndValidate(ctx, &payload) != nil {
		return
	}

	data, code, _, errors := ctl.Repo.SaveMosque(&payload)
	if utils.HandleRepoError(ctx, code, errors) {
		return
	}
	utils.MutationResponse(ctx, code, "Mosque registered successfully", data)
}
