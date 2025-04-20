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

	data, total, page, limit := ctl.Repo.Index(c, filters)
	utils.PaginatedResponse(c, data, total, page, limit)
}

func (ctl *BeneficiaryController) CreateBeneficiary(c *gin.Context) {
	var beneficiary models.Beneficiary

	if utils.BindAndValidate(c, &beneficiary) != nil {
		return
	}

	ctl.Repo.Save(&beneficiary)
	utils.SuccessResponse(c, beneficiary)
}

func (ctl *BeneficiaryController) UpdateBeneficiary(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var beneficiary models.Beneficiary

	if utils.BindAndValidate(c, &beneficiary) != nil {
		return
	}

	updated := ctl.Repo.Update(uint(id), &beneficiary)
	if !updated {
		utils.ErrorResponse(c, http.StatusNotFound, "Data not found")
	}

	utils.SuccessResponse(c, beneficiary)
}

func (ctl *BeneficiaryController) DeleteBeneficiary(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	deleted := ctl.Repo.Delete(uint(id))

	if !deleted {
		utils.ErrorResponse(c, http.StatusNotFound, "Data not found")
		return
	}

	utils.DeleteResponse(c)
}
