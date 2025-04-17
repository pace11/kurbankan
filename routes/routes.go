package routes

import (
	"kurbankan/controllers"
	"kurbankan/repository"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	QurbanPeriodController := controllers.NewQurbanPeriodController(repository.NewQurbanPeriodRepository())

	qurbanPeriod := r.Group("/qurban-periods")
	qurbanPeriod.GET("/", QurbanPeriodController.GetQurbanPeriods)
	qurbanPeriod.POST("/", QurbanPeriodController.CreateQurbanPeriod)
	qurbanPeriod.PATCH("/:id", QurbanPeriodController.UpdateQurbanPeriod)
	qurbanPeriod.DELETE("/:id", QurbanPeriodController.DeleteQurbanPeriod)
}
