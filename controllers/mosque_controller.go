package controllers

import (
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
	mosques := ctl.Repo.Index()
	utils.SuccessResponse(c, mosques)
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
