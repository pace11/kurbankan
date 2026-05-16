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
	data, _, _, total, page, limit := ctl.Repo.Index(c)
	utils.PaginatedResponse(c, data, total, page, limit)
}

func (ctl *TransactionController) GetTransactionsByMosqueID(ctx *gin.Context) {
	// Get mosque member data from context
	mosque, err, code, errors := utils.GetMosqueMemberByContext(ctx)
	if err != nil {
		utils.HandleRepoError(ctx, code, errors)
		return
	}

	data, _, _, total, page, limit := ctl.Repo.IndexByMosqueID(ctx, mosque.UserID)
	utils.PaginatedResponse(ctx, data, total, page, limit)
}

func (ctl *TransactionController) CreateTransaction(ctx *gin.Context) {
	var transaction models.TransactionCreatePayload

	if utils.BindAndValidate(ctx, &transaction) != nil {
		return
	}

	// Get participant data from context
	participant, err, code, errors := utils.GetParticipantByContext(ctx)
	if err != nil {
		utils.HandleRepoError(ctx, code, errors)
		return
	}

	transaction.CreatedByUserID = *participant.UserID

	data, code, entity, errors := ctl.Repo.Create(&transaction)
	if utils.HandleRepoError(ctx, code, errors) {
		return
	}
	utils.MutationResponse(ctx, code, utils.MutationMessage(entity, ctx.Request.Method), data)
}

func (ctl *TransactionController) UploadProof(ctx *gin.Context) {
	var payload models.TransactionUploadProofPayload

	if utils.BindAndValidate(ctx, &payload) != nil {
		return
	}

	// Get transaction ID from URL param
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.ValidationErrorResponse(ctx, map[string]string{"id": "Invalid transaction ID"})
		return
	}
	payload.ID = uint(id)

	data, code, entity, errors := ctl.Repo.UpdateProof(&payload)
	if utils.HandleRepoError(ctx, code, errors) {
		return
	}
	utils.MutationResponse(ctx, code, utils.MutationMessage(entity, ctx.Request.Method), data)
}

func (ctl *TransactionController) VerifyTransaction(ctx *gin.Context) {
	var payload models.TransactionVerifyPayload

	if utils.BindAndValidate(ctx, &payload) != nil {
		return
	}

	// Get transaction ID from URL param
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.ValidationErrorResponse(ctx, map[string]string{"id": "Invalid transaction ID"})
		return
	}

	// Get mosque member data from context
	mosque, err, code, errors := utils.GetMosqueMemberByContext(ctx)
	if err != nil {
		utils.HandleRepoError(ctx, code, errors)
		return
	}

	payload.ID = uint(id)
	payload.VerifiedByUserID = mosque.UserID

	data, code, entity, errors := ctl.Repo.Verify(&payload)
	if utils.HandleRepoError(ctx, code, errors) {
		return
	}
	utils.MutationResponse(ctx, code, utils.MutationMessage(entity, ctx.Request.Method), data)
}

func (ctl *TransactionController) CancelExpiredTransactions(ctx *gin.Context) {
	data, code, entity, errors := ctl.Repo.MarkExpiredPendingTransactions()
	if utils.HandleRepoError(ctx, code, errors) {
		return
	}
	utils.MutationResponse(ctx, code, utils.MutationMessage(entity, ctx.Request.Method), data)
}
