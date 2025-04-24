package controllers

import (
	"kurbankan/repository"
	"kurbankan/utils"

	"github.com/gin-gonic/gin"
)

type RegencyController struct {
	Repo repository.RegencyRepository
}

func NewRegencyRepository(repo repository.RegencyRepository) *RegencyController {
	return &RegencyController{Repo: repo}
}

func (ctl *RegencyController) GetRegencies(c *gin.Context) {
	filters := map[string]any{
		"name":          c.Query("name"),
		"code":          c.Query("code"),
		"province_code": c.Query("province_code"),
	}
	data, code, entity, action, total, page, limit := ctl.Repo.Index(c, filters)
	utils.PaginatedResponse(c, data, code, entity, action, total, page, limit)
}
