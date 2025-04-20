package routes

import (
	"kurbankan/controllers"
	"kurbankan/middlewares"
	"kurbankan/repository"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// controllers
	QurbanPeriodController := controllers.NewQurbanPeriodController(repository.NewQurbanPeriodRepository())
	QurbanOptionController := controllers.NewQurbanOptionController(repository.NewQurbanOptionRepository())
	UserController := controllers.NewUserController(repository.NewUserRepository())
	MosqueController := controllers.NewMosqueRepository(repository.NewMosqueRepository())
	ParticipantController := controllers.NewParticipantRepository(repository.NewParticipantRepository())
	ProvinceController := controllers.NewProvinceRepository(repository.NewProvinceRepository())
	RegencyController := controllers.NewRegencyRepository(repository.NewRegencyRepository())
	DistrictController := controllers.NewDistrictRepository(repository.NewDistrictRepository())
	VillageController := controllers.NewVillageRepository(repository.NewVillageRepository())

	// auth
	auth := r.Group("/auth")
	auth.POST("/register", controllers.RegisterParticipant)
	auth.POST("/register/mosque", controllers.RegisterMosque)
	auth.POST("/login", controllers.Login)

	api := r.Group("/api")

	// area
	area := api.Group("/area")
	area.GET("/provinces", ProvinceController.GetProvinces)
	area.GET("/regencies", RegencyController.GetRegencies)
	area.GET("/districts", DistrictController.GetDistricts)
	area.GET("/villages", VillageController.GetVillages)

	api.Use(middlewares.JWTAuthMiddleware())

	// qurban-periods
	api.GET("/qurban-periods", QurbanPeriodController.GetQurbanPeriods)
	api.POST("/qurban-periods", QurbanPeriodController.CreateQurbanPeriod)
	api.PATCH("/qurban-periods/:id", QurbanPeriodController.UpdateQurbanPeriod)
	api.DELETE("/qurban-periods/:id", QurbanPeriodController.DeleteQurbanPeriod)

	// qurban-options
	api.GET("/qurban-options", QurbanOptionController.GetQurbanOptions)
	api.POST("/qurban-options", QurbanOptionController.CreateQurbanOption)
	api.PATCH("/qurban-options/:id", QurbanOptionController.UpdateQurbanOption)
	api.DELETE("/qurban-options/:id", QurbanOptionController.DeleteQurbanPeriod)

	// users
	api.GET("/users", UserController.GetUsers)
	api.POST("/users", UserController.CreateUser)
	api.PATCH("/users/:id", UserController.UpdateUser)
	api.DELETE("/users/:id", UserController.DeleteUser)

	// mosques
	api.GET("/mosques", MosqueController.GetMosques)
	api.GET("/mosques/:id", MosqueController.GetMosque)
	api.PATCH("/mosques/:id", MosqueController.UpdateMosque)
	api.DELETE("/mosques/:id", MosqueController.DeleteMosque)

	// participants
	api.GET("/participants", ParticipantController.GetParticipants)
	api.GET("/participants/:id", ParticipantController.GetParticipant)

}
