package controllers

import (
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
	participants := ctl.Repo.Index()
	utils.SuccessResponse(c, participants)
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
