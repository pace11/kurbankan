package controllers

import (
	"kurbankan/repository"
	"kurbankan/utils"

	"github.com/gin-gonic/gin"
)

type BankController struct {
	BankRepo repository.BankRepository
}

func NewBankController(bankRepo repository.BankRepository) *BankController {
	return &BankController{
		BankRepo: bankRepo,
	}
}

func (ctl *BankController) GetBanks(c *gin.Context) {
	filters := map[string]any{
		"name": c.Query("name"),
		"code": c.Query("code"),
	}
	data, code, entity, total, page, limit := ctl.BankRepo.Index(c, filters)
	utils.PaginatedResponse(c, data, code, entity, c.Request.Method, total, page, limit)
}
