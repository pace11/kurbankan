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

func NewParticipantRepository(repo repository.ParticipantRepository) *ParticipantController {
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

	data, total, page, limit := ctl.Repo.Index(c, filters)
	utils.PaginatedResponse(c, data, total, page, limit)
}

func (ctl *ParticipantController) GetParticipant(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	participant, err := ctl.Repo.Show(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Participant not found")
		return
	}

	utils.SuccessResponse(c, participant)
}

func (ctl *ParticipantController) UpdateParticipant(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var participant models.UserUpdateDTO

	if utils.BindAndValidate(c, &participant) != nil {
		return
	}

	updated := ctl.Repo.Update(uint(id), &participant)
	if !updated {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to update participant")
		return
	}

	utils.SuccessResponse(c, participant)
}

func (ctl *ParticipantController) DeleteParticipant(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	deleted := ctl.Repo.Delete(uint(id))

	if !deleted {
		utils.ErrorResponse(c, http.StatusFound, "Data not found")
		return
	}

	utils.DeleteResponse(c)
}
