package controllers

import (
	"kurbankan/repository"
	"kurbankan/utils"

	"github.com/gin-gonic/gin"
)

type ProvinceController struct {
	Repo repository.ProvinceRepository
}

func NewProvinceRepository(repo repository.ProvinceRepository) *ProvinceController {
	return &ProvinceController{Repo: repo}
}

func (ctl *ProvinceController) GetProvinces(c *gin.Context) {
	filters := map[string]any{
		"name": c.Query("name"),
		"code": c.Query("code"),
	}
	data, code, entity, action, total, page, limit := ctl.Repo.Index(c, filters)
	utils.PaginatedResponse(c, data, code, entity, action, total, page, limit)
}
