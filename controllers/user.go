package controllers

import (
	"kurbankan/models"
	"kurbankan/repository"
	"kurbankan/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Repo repository.UserRepository
}

func NewUserController(repo repository.UserRepository) *UserController {
	return &UserController{Repo: repo}
}

func (ctl *UserController) GetUsers(c *gin.Context) {
	filters := map[string]any{
		"email": c.Query("email"),
		"role":  c.Query("role"),
	}

	data, code, entity, total, page, limit := ctl.Repo.Index(c, filters)
	utils.PaginatedResponse(c, data, code, entity, c.Request.Method, total, page, limit)
}

func (ctl *UserController) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User

	if utils.BindAndValidate(c, &user) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Update(uint(id), &user)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
