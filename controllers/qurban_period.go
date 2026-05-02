package controllers

import (
	"kurbankan/models"
	"kurbankan/repository"
	"kurbankan/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QurbanPeriodController struct {
	Repo repository.QurbanPeriodRepository
}

func NewQurbanPeriodController(repo repository.QurbanPeriodRepository) *QurbanPeriodController {
	return &QurbanPeriodController{Repo: repo}
}

func (ctl *QurbanPeriodController) GetQurbanPeriods(ctx *gin.Context) {
	filters := map[string]any{
		"year": ctx.Query("year"),
	}

	data, code, entity, total, page, limit := ctl.Repo.Index(ctx, filters)
	utils.PaginatedResponse(ctx, data, code, entity, ctx.Request.Method, total, page, limit)
}

func (ctl *QurbanPeriodController) CreateQurbanPeriod(ctx *gin.Context) {
	var qurbanPeriod models.QurbanPeriod

	if utils.BindAndValidate(ctx, &qurbanPeriod) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Save(&qurbanPeriod)
	utils.HttpResponse(ctx, data, code, entity, ctx.Request.Method, errors)
}

func (ctl *QurbanPeriodController) UpdateQurbanPeriod(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var qurbanPeriod models.QurbanPeriod

	if utils.BindAndValidate(ctx, &qurbanPeriod) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Update(uint(id), &qurbanPeriod)
	utils.HttpResponse(ctx, data, code, entity, ctx.Request.Method, errors)
}

func (ctl *QurbanPeriodController) DeleteQurbanPeriod(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, code, entity, errors := ctl.Repo.Delete(uint(id))
	utils.HttpResponse(ctx, data, code, entity, ctx.Request.Method, errors)
}
