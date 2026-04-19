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

func (ctl *QurbanPeriodController) GetQurbanPeriods(c *gin.Context) {
	filters := map[string]any{
		"year": c.Query("year"),
	}

	data, code, entity, total, page, limit := ctl.Repo.Index(c, filters)
	utils.PaginatedResponse(c, data, code, entity, c.Request.Method, total, page, limit)
}

func (ctl *QurbanPeriodController) CreateQurbanPeriod(c *gin.Context) {
	var qurbanPeriod models.QurbanPeriod

	if utils.BindAndValidate(c, &qurbanPeriod) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Save(&qurbanPeriod)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

func (ctl *QurbanPeriodController) UpdateQurbanPeriod(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var qurbanPeriod models.QurbanPeriod

	if utils.BindAndValidate(c, &qurbanPeriod) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Update(uint(id), &qurbanPeriod)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

func (ctl *QurbanPeriodController) DeleteQurbanPeriod(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code, entity, errors := ctl.Repo.Delete(uint(id))
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
