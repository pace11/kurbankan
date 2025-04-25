package controllers

import (
	"kurbankan/models"
	"kurbankan/repository"
	"kurbankan/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QurbanOptionController struct {
	Repo repository.QurbanOptionRepository
}

func NewQurbanOptionController(repo repository.QurbanOptionRepository) *QurbanOptionController {
	return &QurbanOptionController{Repo: repo}
}

func (ctl *QurbanOptionController) GetQurbanOptions(c *gin.Context) {
	filters := map[string]any{
		"animal_type": c.Query("animal_type"),
		"scheme_type": c.Query("scheme_type"),
	}

	data, code, entity, total, page, limit := ctl.Repo.Index(c, filters)
	utils.PaginatedResponse(c, data, code, entity, c.Request.Method, total, page, limit)
}

func (ctl *QurbanOptionController) CreateQurbanOption(c *gin.Context) {
	var qurbanOption models.QurbanOption

	if utils.BindAndValidate(c, &qurbanOption) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Save(&qurbanOption)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

func (ctl *QurbanOptionController) UpdateQurbanOption(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var qurbanOption models.QurbanOption

	if utils.BindAndValidate(c, &qurbanOption) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Update(uint(id), &qurbanOption)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

func (ctl *QurbanOptionController) DeleteQurbanPeriod(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code, entity, errors := ctl.Repo.Delete(uint(id))
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
