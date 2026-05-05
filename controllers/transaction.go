package controllers

import (
	"kurbankan/models"
	"kurbankan/repository"
	"kurbankan/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	Repo repository.TransactionRepository
}

func NewTransactionController(repo repository.TransactionRepository) *TransactionController {
	return &TransactionController{Repo: repo}
}

func (ctl *TransactionController) GetTransactions(c *gin.Context) {
	data, code, entity, total, page, limit := ctl.Repo.Index(c)
	utils.PaginatedResponse(c, data, code, entity, c.Request.Method, total, page, limit)
}

func (ctl *TransactionController) GetTransactionsByMosqueID(ctx *gin.Context) {
	mosque, err, code, entity, errors := utils.GetMosqueMemberByContext(ctx)
	if err != nil {
		utils.HttpResponse(ctx, nil, code, entity, ctx.Request.Method, errors)
		return
	}

	data, statusCode, entityName, total, page, limit := ctl.Repo.IndexByMosqueID(ctx, mosque.UserID)
	utils.PaginatedResponse(ctx, data, statusCode, entityName, ctx.Request.Method, total, page, limit)
}

func (ctl *TransactionController) CreateTransaction(ctx *gin.Context) {
	var transaction models.TransactionCreatePayload

	if utils.BindAndValidate(ctx, &transaction) != nil {
		return
	}

	// Get user_id from JWT context
	userID := ctx.GetUint("user_id")
	transaction.CreatedByUserID = userID

	data, code, entity, errors := ctl.Repo.Save(&transaction)
	utils.HttpResponse(ctx, data, code, entity, ctx.Request.Method, errors)
}

func (ctl *TransactionController) UploadProof(ctx *gin.Context) {
	var payload models.TransactionUploadProofPayload

	if utils.BindAndValidate(ctx, &payload) != nil {
		return
	}

	// Get transaction ID from URL param
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.HttpResponse(ctx, nil, 400, "transaction", ctx.Request.Method, map[string]string{"id": "Invalid transaction ID"})
		return
	}
	payload.ID = uint(id)

	data, code, entity, errors := ctl.Repo.UpdateProof(&payload)
	utils.HttpResponse(ctx, data, code, entity, ctx.Request.Method, errors)
}
