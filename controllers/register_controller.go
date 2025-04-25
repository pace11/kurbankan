package controllers

import (
	"kurbankan/models"
	"kurbankan/repository"
	"kurbankan/utils"

	"github.com/gin-gonic/gin"
)

type RegisterController struct {
	Repo repository.RegisterRepository
}

func NewRegisterController(repo repository.RegisterRepository) *RegisterController {
	return &RegisterController{Repo: repo}
}

func (ctl *RegisterController) RegisterParticipant(c *gin.Context) {
	var payload models.UserCreateDTO

	if utils.BindAndValidate(c, &payload) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.SaveParticipant(&payload)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

func (ctl *RegisterController) RegisterMosque(c *gin.Context) {
	var payload models.UserCreateDTO

	if utils.BindAndValidate(c, &payload) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.SaveMosque(&payload)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
