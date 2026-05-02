package controllers

import (
	"kurbankan/models"
	"kurbankan/repository"
	"kurbankan/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QurbanOfferingController struct {
	Repo repository.QurbanOfferingRepository
}

func NewQurbanOfferingController(repo repository.QurbanOfferingRepository) *QurbanOfferingController {
	return &QurbanOfferingController{Repo: repo}
}

func (ctl *QurbanOfferingController) GetQurbanOfferings(ctx *gin.Context) {
	filters := map[string]any{
		"animal_type": ctx.Query("animal_type"),
		"scheme_type": ctx.Query("scheme_type"),
	}

	data, code, entity, total, page, limit := ctl.Repo.Index(ctx, filters)
	utils.PaginatedResponse(ctx, data, code, entity, ctx.Request.Method, total, page, limit)
}

func (ctl *QurbanOfferingController) CreateQurbanOffering(ctx *gin.Context) {
	var qurbanOffering models.QurbanOffering

	if utils.BindAndValidate(ctx, &qurbanOffering) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Save(&qurbanOffering)
	utils.HttpResponse(ctx, data, code, entity, ctx.Request.Method, errors)
}

func (ctl *QurbanOfferingController) UpdateQurbanOffering(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var qurbanOffering models.QurbanOffering

	if utils.BindAndValidate(ctx, &qurbanOffering) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Update(uint(id), &qurbanOffering)
	utils.HttpResponse(ctx, data, code, entity, ctx.Request.Method, errors)
}

func (ctl *QurbanOfferingController) DeleteQurbanOffering(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, code, entity, errors := ctl.Repo.Delete(uint(id))
	utils.HttpResponse(ctx, data, code, entity, ctx.Request.Method, errors)
}
