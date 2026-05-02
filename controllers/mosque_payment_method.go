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

	mosque, err, code, entity, errors := utils.GetMosqueMemberByContext(ctx)
	if err != nil {
		utils.HttpResponse(ctx, nil, code, entity, ctx.Request.Method, errors)
		return
	}

	payload.MosqueID = mosque.ID

	data, code, entity, errors := ctl.Repo.Create(&payload)
	utils.HttpResponse(ctx, data, code, entity, ctx.Request.Method, errors)
}
