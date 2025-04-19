package controllers

import (
	"kurbankan/repository"
	"kurbankan/utils"

	"github.com/gin-gonic/gin"
)

type DistrictController struct {
	Repo repository.DistrictRepository
}

func NewDistrictRepository(repo repository.DistrictRepository) *DistrictController {
	return &DistrictController{Repo: repo}
}

func (ctl *DistrictController) GetDistricts(c *gin.Context) {
	filters := map[string]any{
		"name":         c.Query("name"),
		"code":         c.Query("code"),
		"regency_code": c.Query("regency_code"),
	}
	data, total, page, limit := ctl.Repo.Index(c, filters)
	utils.PaginatedResponse(c, data, total, page, limit)
}
