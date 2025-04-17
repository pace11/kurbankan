package controllers

import (
	"kurbankan/models"
	"kurbankan/repository"
	"kurbankan/utils"
	"net/http"
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
	users := ctl.Repo.Index()
	utils.SuccessResponse(c, users)
}

func (ctl *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if utils.BindAndValidate(c, &user) != nil {
		return
	}
	ctl.Repo.Save(&user)
	utils.SuccessResponse(c, user)
}

func (ctl *UserController) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	if utils.BindAndValidate(c, &user) != nil {
		return
	}
	updated := ctl.Repo.Update(uint(id), &user)
	if !updated {
		utils.ErrorResponse(c, http.StatusNotFound, "Data not found")
	}
	utils.SuccessResponse(c, user)
}

func (ctl *UserController) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	deleted := ctl.Repo.Delete(uint(id))
	if !deleted {
		utils.ErrorResponse(c, http.StatusNotFound, "Data not found")
		return
	}
	utils.DeleteResponse(c)
}
