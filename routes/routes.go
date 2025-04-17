package routes

import (
	"kurbankan/controllers"
	"kurbankan/repository"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	QurbanPeriodController := controllers.NewQurbanPeriodController(repository.NewQurbanPeriodRepository())
	UserController := controllers.NewUserController(repository.NewUserRepository())

	api := r.Group("/api")
	// auth
	auth := api.Group("/auth")
	auth.POST("/register", controllers.RegisterParticipant)
	auth.POST("/register/mosque", controllers.RegisterMosque)

	// qurban-periods endpoint
	api.GET("/qurban-periods", QurbanPeriodController.GetQurbanPeriods)
	api.POST("/qurban-periods", QurbanPeriodController.CreateQurbanPeriod)
	api.PATCH("/qurban-periods/:id", QurbanPeriodController.UpdateQurbanPeriod)
	api.DELETE("/qurban-periods/:id", QurbanPeriodController.DeleteQurbanPeriod)

	// users
	api.GET("/users", UserController.GetUsers)
	api.POST("/users", UserController.CreateUser)
	api.PATCH("/users/:id", UserController.UpdateUser)
	api.DELETE("/users/:id", UserController.DeleteUser)

}
