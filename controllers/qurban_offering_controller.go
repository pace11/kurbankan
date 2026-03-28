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

func (ctl *QurbanOfferingController) GetQurbanOfferings(c *gin.Context) {
	filters := map[string]any{
		"animal_type": c.Query("animal_type"),
		"scheme_type": c.Query("scheme_type"),
	}

	data, code, entity, total, page, limit := ctl.Repo.Index(c, filters)
	utils.PaginatedResponse(c, data, code, entity, c.Request.Method, total, page, limit)
}

func (ctl *QurbanOfferingController) CreateQurbanOffering(c *gin.Context) {
	var qurbanOffering models.QurbanOffering

	if utils.BindAndValidate(c, &qurbanOffering) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Save(&qurbanOffering)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

func (ctl *QurbanOfferingController) UpdateQurbanOffering(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var qurbanOffering models.QurbanOffering

	if utils.BindAndValidate(c, &qurbanOffering) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Update(uint(id), &qurbanOffering)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

func (ctl *QurbanOfferingController) DeleteQurbanOffering(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code, entity, errors := ctl.Repo.Delete(uint(id))
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
