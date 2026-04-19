package controllers

import (
	"kurbankan/models"
	"kurbankan/repository"
	"kurbankan/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BeneficiaryController struct {
	Repo repository.BeneficiaryRepository
}

func NewBeneficiaryController(repo repository.BeneficiaryRepository) *BeneficiaryController {
	return &BeneficiaryController{Repo: repo}
}

func (ctl *BeneficiaryController) GetBeneficiaries(c *gin.Context) {
	filters := map[string]any{
		"name": c.Query("name"),
	}

	data, code, entity, total, page, limit := ctl.Repo.Index(c, filters)
	utils.PaginatedResponse(c, data, code, entity, c.Request.Method, total, page, limit)
}

func (ctl *BeneficiaryController) GetBeneficiary(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		utils.HttpResponse(c, nil, http.StatusBadRequest, "Invalid ID", c.Request.Method, nil)
		return
	}

	data, code, entity, errors := ctl.Repo.Show(uint(id))
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

func (ctl *BeneficiaryController) CreateBeneficiary(c *gin.Context) {
	var beneficiary models.Beneficiary

	if utils.BindAndValidate(c, &beneficiary) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Save(&beneficiary)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

func (ctl *BeneficiaryController) UpdateBeneficiary(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var beneficiary models.Beneficiary

	if utils.BindAndValidate(c, &beneficiary) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Update(uint(id), &beneficiary)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

func (ctl *BeneficiaryController) DeleteBeneficiary(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code, entity, errors := ctl.Repo.Delete(uint(id))
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
