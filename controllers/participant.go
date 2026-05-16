package controllers

import (
	"kurbankan/models"
	"kurbankan/repository"
	"kurbankan/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ParticipantController struct {
	Repo repository.ParticipantRepository
}

func NewParticipantController(repo repository.ParticipantRepository) *ParticipantController {
	return &ParticipantController{Repo: repo}
}

func (ctl *ParticipantController) GetParticipants(c *gin.Context) {
	filters := map[string]any{
		"name":          c.Query("name"),
		"province_code": c.Query("province_code"),
		"regency_code":  c.Query("regency_code"),
		"district_code": c.Query("district_code"),
		"village_code":  c.Query("village_code"),
	}

	data, _, _, total, page, limit := ctl.Repo.Index(c, filters)
	utils.PaginatedResponse(c, data, total, page, limit)
}

func (ctl *ParticipantController) GetParticipant(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "BAD_REQUEST", "Invalid ID")
		return
	}

	data, code, _, errors := ctl.Repo.Show(uint(id))
	if utils.HandleRepoError(c, code, errors) {
		return
	}
	utils.DetailResponse(c, data)
}

func (ctl *ParticipantController) UpdateParticipant(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var participant models.UserUpdatePayload

	if utils.BindAndValidate(c, &participant) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Update(uint(id), &participant)
	if utils.HandleRepoError(c, code, errors) {
		return
	}
	utils.MutationResponse(c, code, utils.MutationMessage(entity, c.Request.Method), data)
}

func (ctl *ParticipantController) DeleteParticipant(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code, entity, errors := ctl.Repo.Delete(uint(id))
	if utils.HandleRepoError(c, code, errors) {
		return
	}
	utils.MutationResponse(c, code, utils.MutationMessage(entity, c.Request.Method), data)
}
