package controllers

import (
	"kurbankan/repository"
	"kurbankan/utils"

	"github.com/gin-gonic/gin"
)

type VillageController struct {
	Repo repository.VillageRepository
}

func NewVillageRepository(repo repository.VillageRepository) *VillageController {
	return &VillageController{Repo: repo}
}

func (ctl *VillageController) GetVillages(c *gin.Context) {
	filters := map[string]any{
		"name":          c.Query("name"),
		"code":          c.Query("code"),
		"district_code": c.Query("district_code"),
	}
	data, code, entity, action, total, page, limit := ctl.Repo.Index(c, filters)
	utils.PaginatedResponse(c, data, code, entity, action, total, page, limit)
}
