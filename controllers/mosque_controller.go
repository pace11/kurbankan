package controllers

import (
	"kurbankan/models"
	"kurbankan/repository"
	"kurbankan/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MosqueController struct {
	Repo repository.MosqueRepository
}

func NewMosqueRepository(repo repository.MosqueRepository) *MosqueController {
	return &MosqueController{Repo: repo}
}

func (ctl *MosqueController) GetMosques(c *gin.Context) {
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

func (ctl *MosqueController) GetMosque(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		utils.HttpResponse(c, nil, http.StatusBadRequest, "Invalid ID", c.Request.Method, nil)
		return
	}

	data, code, entity, errors := ctl.Repo.Show(uint(id))
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

func (ctl *MosqueController) CreateMosque(c *gin.Context) {
	var mosque models.UserCreateDTO

	if utils.BindAndValidate(c, &mosque) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Save(&mosque)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

func (ctl *MosqueController) UpdateMosque(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var mosque models.UserUpdateDTO

	if utils.BindAndValidate(c, &mosque) != nil {
		return
	}

	data, code, entity, errors := ctl.Repo.Update(uint(id), &mosque)
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}

func (ctl *MosqueController) DeleteMosque(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code, entity, errors := ctl.Repo.Delete(uint(id))
	utils.HttpResponse(c, data, code, entity, c.Request.Method, errors)
}
