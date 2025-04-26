package controllers

import (
	"kurbankan/models"
	"kurbankan/repository"
	"kurbankan/utils"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	Repo repository.LoginRepository
}

func NewLoginController(repo repository.LoginRepository) *LoginController {
	return &LoginController{Repo: repo}
}

func (ctl *LoginController) Login(c *gin.Context) {
	var payload models.LoginDTO

	if utils.BindAndValidate(c, &payload) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Login(&payload)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
