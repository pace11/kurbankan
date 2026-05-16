package controllers

import (
	"kurbankan/models"
	"kurbankan/repository"
	"kurbankan/utils"

	"github.com/gin-gonic/gin"
)

type MosquePaymentMethodController struct {
	Repo repository.MosquePaymentMethodRepository
}

func NewMosquePaymentMethodController(repo repository.MosquePaymentMethodRepository) *MosquePaymentMethodController {
	return &MosquePaymentMethodController{Repo: repo}
}

func (ctl *MosquePaymentMethodController) Create(ctx *gin.Context) {
	var payload models.MosquePaymentMethodCreatePayload

	if utils.BindAndValidate(ctx, &payload) != nil {
		return
	}

	mosque, err, code, errors := utils.GetMosqueMemberByContext(ctx)
	if err != nil {
		utils.HandleRepoError(ctx, code, errors)
		return
	}

	payload.MosqueID = mosque.ID

	data, code, entity, errors := ctl.Repo.Create(&payload)
	if utils.HandleRepoError(ctx, code, errors) {
		return
	}
	utils.MutationResponse(ctx, code, utils.MutationMessage(entity, ctx.Request.Method), data)
}
