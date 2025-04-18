package routes

import (
	"kurbankan/controllers"
	"kurbankan/middlewares"
	"kurbankan/repository"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	QurbanPeriodController := controllers.NewQurbanPeriodController(repository.NewQurbanPeriodRepository())
	UserController := controllers.NewUserController(repository.NewUserRepository())
	MosqueController := controllers.NewMosqueRepository(repository.NewMosqueRepository())

	// auth
	auth := r.Group("/auth")
	auth.POST("/register", controllers.RegisterParticipant)
	auth.POST("/register/mosque", controllers.RegisterMosque)
	auth.POST("/login", controllers.Login)

	api := r.Group("/api")
	api.Use(middlewares.JWTAuthMiddleware())

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

	// mosques
	api.GET("/mosques", MosqueController.GetMosques)
	api.GET("/mosques/:id", MosqueController.GetMosque)

}
