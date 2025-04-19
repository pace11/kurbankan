package controllers

import (
	"kurbankan/models"
	"kurbankan/repository"
	"kurbankan/utils"
	"net/http"
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
	qurbanOptions := ctl.Repo.Index()
	utils.SuccessResponse(c, qurbanOptions)
}

func (ctl *QurbanOptionController) CreateQurbanOption(c *gin.Context) {
	var qurbanOption models.QurbanOption
	if utils.BindAndValidate(c, &qurbanOption) != nil {
		return
	}
	ctl.Repo.Save(&qurbanOption)
	utils.SuccessResponse(c, qurbanOption)
}

func (ctl *QurbanOptionController) UpdateQurbanOption(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var qurbanOption models.QurbanOption
	if utils.BindAndValidate(c, &qurbanOption) != nil {
		return
	}
	updated := ctl.Repo.Update(uint(id), &qurbanOption)
	if !updated {
		utils.ErrorResponse(c, http.StatusNotFound, "Data not found")
	}
	utils.SuccessResponse(c, qurbanOption)
}

func (ctl *QurbanOptionController) DeleteQurbanPeriod(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	deleted := ctl.Repo.Delete(uint(id))
	if !deleted {
		utils.ErrorResponse(c, http.StatusNotFound, "Data not found")
		return
	}
	utils.DeleteResponse(c)
}
