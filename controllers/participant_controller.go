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

	data, code, entity, total, page, limit := ctl.Repo.Index(c, filters)
	utils.PaginatedResponse(c, data, code, entity, c.Request.Method, total, page, limit)
}

func (ctl *ParticipantController) GetParticipant(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		utils.HttpResponse(c, nil, http.StatusBadRequest, "Invalid ID", c.Request.Method, nil)
		return
	}

	data, code, entity, errors := ctl.Repo.Show(uint(id))
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

func (ctl *ParticipantController) UpdateParticipant(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var participant models.UserUpdateDTO

	if utils.BindAndValidate(c, &participant) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Update(uint(id), &participant)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

func (ctl *ParticipantController) DeleteParticipant(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code, entity, errors := ctl.Repo.Delete(uint(id))
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
