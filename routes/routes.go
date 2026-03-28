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
	QurbanOfferingController := controllers.NewQurbanOfferingController(repository.NewQurbanOfferingRepository())
	UserController := controllers.NewUserController(repository.NewUserRepository())
	MosqueController := controllers.NewMosqueController(repository.NewMosqueRepository())
	ParticipantController := controllers.NewParticipantController(repository.NewParticipantRepository())
	ProvinceController := controllers.NewProvinceController(repository.NewProvinceRepository())
	RegencyController := controllers.NewRegencyController(repository.NewRegencyRepository())
	DistrictController := controllers.NewDistrictController(repository.NewDistrictRepository())
	VillageController := controllers.NewVillageController(repository.NewVillageRepository())
	BeneficiaryController := controllers.NewBeneficiaryController(repository.NewBeneficiaryRepository())
	RegisterController := controllers.NewRegisterController(repository.NewRegisterRepository())
	LoginController := controllers.NewLoginController(repository.NewLoginRepository())
	TransactionController := controllers.NewTransactionController(repository.NewTransactionRepository())
	MigrationController := controllers.NewMigrationController()
	SeederController := controllers.NewSeederController()

	// auth
	auth := r.Group("/auth")
	auth.POST("/register/participant", RegisterController.RegisterParticipant)
	auth.POST("/register/mosque", RegisterController.RegisterMosque)
	auth.POST("/login", LoginController.Login)

	api := r.Group("/api")

	// migration endpoints
	api.GET("/migrate/status", MigrationController.CheckMigrationStatus)
	api.POST("/migrate/up", MigrationController.RunMigration)
	api.POST("/migrate/down", MigrationController.DropAllTables)

	// seeder endpoints
	api.GET("/seed/status", SeederController.GetSeederStatus)
	api.POST("/seed/run", SeederController.RunAllSeeders)
	api.POST("/seed/run/:seed", SeederController.RunSpecificSeeder)
	api.DELETE("/seed/clear", SeederController.ClearSeededData)

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

	// qurban-offerings
	api.GET("/qurban-offerings", QurbanOfferingController.GetQurbanOfferings)
	api.POST("/qurban-offering", QurbanOfferingController.CreateQurbanOffering)
	api.PATCH("/qurban-offering/:id", QurbanOfferingController.UpdateQurbanOffering)
	api.DELETE("/qurban-offering/:id", QurbanOfferingController.DeleteQurbanOffering)

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
	api.GET("/transactions", TransactionController.GetTransactions)
	api.POST("/transaction", TransactionController.CreateTransaction)
}
