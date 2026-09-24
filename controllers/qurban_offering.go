package controllers

import (
	"kurbankan/models"
	"kurbankan/repository"
	"kurbankan/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QurbanOfferingController struct {
	Repo repository.QurbanOfferingRepository
}

func NewQurbanOfferingController(repo repository.QurbanOfferingRepository) *QurbanOfferingController {
	return &QurbanOfferingController{Repo: repo}
}

// GetQurbanOfferingsWithPagination godoc
// @Summary      List qurban offerings
// @Description  Returns a paginated list of qurban offerings. Optionally filter by animal type and scheme type.
// @Tags         Qurban Offerings
// @Produce      json
// @Security     BearerAuth
// @Param        animal_type query     string  false  "Filter by animal type (e.g. cow, goat)"
// @Param        scheme_type query     string  false  "Filter by scheme type (e.g. regular, premium)"
// @Param        page   query     int     false  "Page number"          default(1)
// @Param        limit  query     int     false  "Items per page"       default(10)
// @Success      200    {object}  models.QurbanOfferingListWithPaginationResponse
// @Failure      401    {object}  models.SwaggerErrorResponse
// @Router       /qurban-offerings [get]
func (ctl *QurbanOfferingController) GetQurbanOfferingsWithPagination(ctx *gin.Context) {
	filters := map[string]any{
		"animal_type": ctx.Query("animal_type"),
		"scheme_type": ctx.Query("scheme_type"),
	}

	// Get mosque member data from context
	mosque, err, code, errors := utils.GetMosqueMemberByContext(ctx)
	if err != nil {
		utils.HandleRepoError(ctx, code, errors)
		return
	}

	data, total, page, limit := ctl.Repo.ListWithPagination(ctx, mosque.MosqueID, filters)
	utils.PaginatedResponse(ctx, data, total, page, limit)
}

// CreateQurbanOffering godoc
// @Summary      Create a qurban offering
// @Description  Creates a new qurban offering for the authenticated mosque.
// @Tags         Qurban Offerings
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      models.QurbanOfferingRequest          true  "Qurban offering payload"
// @Success      201   {object}  models.QurbanOfferingMutationResponse
// @Failure      400   {object}  models.SwaggerValidationErrorResponse
// @Failure      401   {object}  models.SwaggerErrorResponse
// @Failure      404   {object}  models.SwaggerErrorResponse
// @Failure      500   {object}  models.SwaggerErrorResponse
// @Router       /qurban-offerings [post]
func (ctl *QurbanOfferingController) CreateQurbanOffering(ctx *gin.Context) {
	var payload models.QurbanOfferingRequest

	if utils.BindAndValidate(ctx, &payload) != nil {
		return
	}

	// Get mosque member data from context
	mosque, err, code, errors := utils.GetMosqueMemberByContext(ctx)
	if err != nil {
		utils.HandleRepoError(ctx, code, errors)
		return
	}

	payload.MosqueID = mosque.MosqueID
	data, code, entity, errors := ctl.Repo.Save(&payload)
	if utils.HandleRepoError(ctx, code, errors) {
		return
	}

	utils.MutationResponse(ctx, code, utils.MutationMessage(entity, ctx.Request.Method), data)
}

// UpdateQurbanOffering godoc
// @Summary      Update a qurban offering
// @Description  Updates an existing qurban offering by ID for the authenticated mosque.
// @Tags         Qurban Offerings
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int                                 true  "Qurban Offering ID"
// @Param        body  body      models.QurbanOfferingRequest          true  "Qurban offering payload"
// @Success      200   {object}  models.QurbanOfferingMutationResponse
// @Failure      400   {object}  models.SwaggerValidationErrorResponse
// @Failure      401   {object}  models.SwaggerErrorResponse
// @Failure      404   {object}  models.SwaggerErrorResponse
// @Failure      500   {object}  models.SwaggerErrorResponse
// @Router       /qurban-offerings/{id} [patch]
func (ctl *QurbanOfferingController) UpdateQurbanOffering(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var payload models.QurbanOfferingRequest

	if utils.BindAndValidate(ctx, &payload) != nil {
		return
	}

	payload.ID = uint(id)
	data, code, entity, errors := ctl.Repo.Update(&payload)
	if utils.HandleRepoError(ctx, code, errors) {
		return
	}

	utils.MutationResponse(ctx, code, utils.MutationMessage(entity, ctx.Request.Method), data)
}

// DeleteQurbanOffering godoc
// @Summary      Delete a qurban offering
// @Description  Deletes a qurban offering by ID for the authenticated mosque.
// @Tags         Qurban Offerings
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      int                                 true  "Qurban Offering ID"
// @Success      200 {object}  models.QurbanOfferingMutationResponse
// @Failure      401 {object}  models.SwaggerErrorResponse
// @Failure      404 {object}  models.SwaggerErrorResponse
// @Failure      500 {object}  models.SwaggerErrorResponse
// @Router       /qurban-offerings/{id} [delete]
func (ctl *QurbanOfferingController) DeleteQurbanOffering(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, code, entity, errors := ctl.Repo.Delete(uint(id))
	if utils.HandleRepoError(ctx, code, errors) {
		return
	}

	utils.MutationResponse(ctx, code, utils.MutationMessage(entity, ctx.Request.Method), data)
}
