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
	var payload models.LoginPayload

	if utils.BindAndValidate(c, &payload) != nil {
		return
	}

	data, code, _, errors := ctl.Repo.Login(&payload)
	if utils.HandleRepoError(c, code, errors) {
		return
	}
	utils.MutationResponse(c, code, "Login successful", data)
}
