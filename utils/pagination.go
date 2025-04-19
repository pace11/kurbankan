package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ApplyPagination(c *gin.Context, db *gorm.DB) (*gorm.DB, int, int) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page < 1 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	return db.Offset(offset).Limit(limit), page, limit
}
