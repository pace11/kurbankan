package controllers

import (
	"kurbankan/models"
	"kurbankan/repository"
	"kurbankan/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QurbanPeriodController struct {
	Repo repository.QurbanPeriodRepository
}

func NewQurbanPeriodController(repo repository.QurbanPeriodRepository) *QurbanPeriodController {
	return &QurbanPeriodController{Repo: repo}
}

// GetQurbanPeriods godoc
// @Summary      List qurban periods
// @Description  Returns a paginated list of qurban periods. Optionally filter by year.
// @Tags         Qurban Periods
// @Produce      json
// @Security     BearerAuth
// @Param        year   query     string  false  "Filter by year (e.g. 2025)"
// @Param        page   query     int     false  "Page number"          default(1)
// @Param        limit  query     int     false  "Items per page"       default(10)
// @Success      200    {object}  models.QurbanPeriodListResponse
// @Failure      401    {object}  models.SwaggerErrorResponse
// @Router       /qurban-periods [get]
func (ctl *QurbanPeriodController) GetQurbanPeriods(ctx *gin.Context) {
	filters := map[string]any{
		"year": ctx.Query("year"),
	}

	data, _, _, total, page, limit := ctl.Repo.Index(ctx, filters)
	utils.PaginatedResponse(ctx, data, total, page, limit)
}

// CreateQurbanPeriod godoc
// @Summary      Create a qurban period
// @Description  Creates a new qurban period for the authenticated mosque.
// @Tags         Qurban Periods
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      models.QurbanPeriodRequest          true  "Qurban period payload"
// @Success      201   {object}  models.QurbanPeriodMutationResponse
// @Failure      400   {object}  models.SwaggerValidationErrorResponse
// @Failure      401   {object}  models.SwaggerErrorResponse
// @Failure      404   {object}  models.SwaggerErrorResponse
// @Failure      500   {object}  models.SwaggerErrorResponse
// @Router       /qurban-periods [post]
func (ctl *QurbanPeriodController) CreateQurbanPeriod(ctx *gin.Context) {
	var payload models.QurbanPeriodRequest

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

// UpdateQurbanPeriod godoc
// @Summary      Update a qurban period
// @Description  Updates an existing qurban period by ID for the authenticated mosque.
// @Tags         Qurban Periods
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int                                 true  "Qurban Period ID"
// @Param        body  body      models.QurbanPeriodRequest          true  "Qurban period payload"
// @Success      200   {object}  models.QurbanPeriodMutationResponse
// @Failure      400   {object}  models.SwaggerValidationErrorResponse
// @Failure      401   {object}  models.SwaggerErrorResponse
// @Failure      404   {object}  models.SwaggerErrorResponse
// @Failure      500   {object}  models.SwaggerErrorResponse
// @Router       /qurban-periods/{id} [patch]
func (ctl *QurbanPeriodController) UpdateQurbanPeriod(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var payload models.QurbanPeriodRequest

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
	payload.ID = uint(id)
	data, code, entity, errors := ctl.Repo.Update(&payload)
	if utils.HandleRepoError(ctx, code, errors) {
		return
	}
	utils.MutationResponse(ctx, code, utils.MutationMessage(entity, ctx.Request.Method), data)
}

// DeleteQurbanPeriod godoc
// @Summary      Delete a qurban period
// @Description  Deletes a qurban period by ID for the authenticated mosque.
// @Tags         Qurban Periods
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      int                                 true  "Qurban Period ID"
// @Success      200 {object}  models.QurbanPeriodMutationResponse
// @Failure      401 {object}  models.SwaggerErrorResponse
// @Failure      404 {object}  models.SwaggerErrorResponse
// @Failure      500 {object}  models.SwaggerErrorResponse
// @Router       /qurban-periods/{id} [delete]
func (ctl *QurbanPeriodController) DeleteQurbanPeriod(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	// Get mosque member data from context
	_, err, code, errors := utils.GetMosqueMemberByContext(ctx)
	if err != nil {
		utils.HandleRepoError(ctx, code, errors)
		return
	}

	data, code, entity, errors := ctl.Repo.Delete(uint(id))
	if utils.HandleRepoError(ctx, code, errors) {
		return
	}
	utils.MutationResponse(ctx, code, utils.MutationMessage(entity, ctx.Request.Method), data)
}
