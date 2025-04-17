package controllers

import (
	"kurbankan/models"
	"kurbankan/repository"
	"kurbankan/utils"
	"net/http"
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
	qurbanPeriods := ctl.Repo.Index()
	utils.SuccessResponse(c, qurbanPeriods)
}

func (ctl *QurbanPeriodController) CreateQurbanPeriod(c *gin.Context) {
	var qurbanPeriod models.QurbanPeriod
	if utils.BindAndValidate(c, &qurbanPeriod) != nil {
		return
	}
	ctl.Repo.Save(&qurbanPeriod)
	utils.SuccessResponse(c, qurbanPeriod)
}

func (ctl *QurbanPeriodController) UpdateQurbanPeriod(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var qurbanPeriod models.QurbanPeriod
	if utils.BindAndValidate(c, &qurbanPeriod) != nil {
		return
	}
	updated := ctl.Repo.Update(uint(id), &qurbanPeriod)
	if !updated {
		utils.ErrorResponse(c, http.StatusNotFound, "Data not found")
	}
	utils.SuccessResponse(c, qurbanPeriod)
}

func (ctl *QurbanPeriodController) DeleteQurbanPeriod(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	deleted := ctl.Repo.Delete(uint(id))
	if !deleted {
		utils.ErrorResponse(c, http.StatusNotFound, "Data not found")
		return
	}
	utils.DeleteResponse(c)
}
