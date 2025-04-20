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

	data, total, page, limit := ctl.Repo.Index(c, filters)
	utils.PaginatedResponse(c, data, total, page, limit)
}

func (ctl *MosqueController) GetMosque(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	mosque, err := ctl.Repo.Show(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Mosque not found")
		return
	}

	utils.SuccessResponse(c, mosque)
}

func (ctl *MosqueController) CreateMosque(c *gin.Context) {
	var mosque models.UserCreateDTO

	if utils.BindAndValidate(c, &mosque) != nil {
		return
	}

	ctl.Repo.Save(&mosque)
	utils.SuccessResponse(c, mosque)
}

func (ctl *MosqueController) UpdateMosque(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var mosque models.UserUpdateDTO

	if utils.BindAndValidate(c, &mosque) != nil {
		return
	}

	updated := ctl.Repo.Update(uint(id), &mosque)
	if !updated {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to update mosque")
		return
	}

	utils.SuccessResponse(c, mosque)
}

func (ctl *MosqueController) DeleteMosque(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	deleted := ctl.Repo.Delete(uint(id))
	if !deleted {
		utils.ErrorResponse(c, http.StatusNotFound, "Data not found")
		return
	}
	utils.DeleteResponse(c)
}
