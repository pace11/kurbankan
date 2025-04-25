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
	BeneficiaryController := controllers.NewBeneficiaryController(repository.NewBeneficiaryRepository())

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

	// xendit virtual account
	api.POST("/xendit/va-callback", controllers.XenditVAWebhookHandler)

	api.Use(middlewares.JWTAuthMiddleware())

	// qurban-periods
	api.GET("/qurban-periods", QurbanPeriodController.GetQurbanPeriods)
	api.POST("/qurban-period", QurbanPeriodController.CreateQurbanPeriod)
	api.PATCH("/qurban-period/:id", QurbanPeriodController.UpdateQurbanPeriod)
	api.DELETE("/qurban-period/:id", QurbanPeriodController.DeleteQurbanPeriod)

	// qurban-options
	api.GET("/qurban-options", QurbanOptionController.GetQurbanOptions)
	api.POST("/qurban-option", QurbanOptionController.CreateQurbanOption)
	api.PATCH("/qurban-option/:id", QurbanOptionController.UpdateQurbanOption)
	api.DELETE("/qurban-option/:id", QurbanOptionController.DeleteQurbanPeriod)

	// users
	api.GET("/users", UserController.GetUsers)
	api.PATCH("/user/:id", UserController.UpdateUser)

	// mosques
	api.GET("/mosques", MosqueController.GetMosques)
	api.GET("/mosque/:id", MosqueController.GetMosque)
	api.POST("/mosque", MosqueController.CreateMosque)
	api.PATCH("/mosque/:id", MosqueController.UpdateMosque)
	api.DELETE("/mosque/:id", MosqueController.DeleteMosque)

	// participants
	api.GET("/participants", ParticipantController.GetParticipants)
	api.GET("/participant/:id", ParticipantController.GetParticipant)
	api.PATCH("/participant/:id", ParticipantController.UpdateParticipant)
	api.DELETE("/participant/:id", ParticipantController.DeleteParticipant)

	// beneficiaries
	api.GET("/beneficiaries", BeneficiaryController.GetBeneficiaries)
	api.GET("/beneficiary/:id", BeneficiaryController.GetBeneficiary)
	api.POST("/beneficiary", BeneficiaryController.CreateBeneficiary)
	api.PATCH("/beneficiary/:id", BeneficiaryController.UpdateBeneficiary)
	api.DELETE("/beneficiary/:id", BeneficiaryController.DeleteBeneficiary)

	// transactions
	api.POST("/transaction", controllers.CreateTransaction)
}
