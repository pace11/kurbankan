package routes

import (
	"kurbankan/controllers"
	"kurbankan/middlewares"
	"kurbankan/repository"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine) {

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

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
	BankController := controllers.NewBankController(repository.NewBankRepository())
	BeneficiaryController := controllers.NewBeneficiaryController(repository.NewBeneficiaryRepository())
	RegisterController := controllers.NewRegisterController(repository.NewRegisterRepository())
	LoginController := controllers.NewLoginController(repository.NewLoginRepository())
	TransactionController := controllers.NewTransactionController(repository.NewTransactionRepository())
	MosquePaymentMethodController := controllers.NewMosquePaymentMethodController(repository.NewMosquePaymentMethodRepository())
	MigrationController := controllers.NewMigrationController()
	SeederController := controllers.NewSeederController()

	// auth
	auth := r.Group("/auth")
	auth.POST("/register/participant", RegisterController.RegisterParticipant)
	auth.POST("/register/mosque", RegisterController.RegisterMosque)
	auth.POST("/login", LoginController.Login)

	// swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

	// constant data
	api.GET("/constants/banks", BankController.GetBanks)

	// xendit virtual account
	api.POST("/xendit/va-callback", controllers.XenditVAWebhookHandler)

	// public endpoints
	api.POST("/transactions/cancel-expired", TransactionController.CancelExpiredTransactions)

	api.Use(middlewares.JWTAuthMiddleware())

	// qurban-periods
	api.GET("/qurban-periods", QurbanPeriodController.GetQurbanPeriods)
	api.POST("/qurban-periods", QurbanPeriodController.CreateQurbanPeriod)
	api.PATCH("/qurban-periods/:id", QurbanPeriodController.UpdateQurbanPeriod)
	api.DELETE("/qurban-periods/:id", QurbanPeriodController.DeleteQurbanPeriod)

	// qurban-offerings
	api.GET("/qurban-offerings", QurbanOfferingController.GetQurbanOfferings)
	api.POST("/qurban-offerings", QurbanOfferingController.CreateQurbanOffering)
	api.PATCH("/qurban-offerings/:id", QurbanOfferingController.UpdateQurbanOffering)
	api.DELETE("/qurban-offerings/:id", QurbanOfferingController.DeleteQurbanOffering)

	// users
	api.GET("/users", UserController.GetUsers)
	api.PATCH("/users/:id", UserController.UpdateUser)

	// mosques
	api.GET("/mosques", MosqueController.GetMosques)
	api.GET("/mosques/:id", MosqueController.GetMosque)
	api.POST("/mosques", MosqueController.CreateMosque)
	api.PATCH("/mosques/:id", MosqueController.UpdateMosque)
	api.DELETE("/mosques/:id", MosqueController.DeleteMosque)

	// mosque payment methods
	api.POST("/mosques/payment-methods", MosquePaymentMethodController.Create)

	// participants
	api.GET("/participants", ParticipantController.GetParticipants)
	api.GET("/participants/:id", ParticipantController.GetParticipant)
	api.PATCH("/participants/:id", ParticipantController.UpdateParticipant)
	api.DELETE("/participants/:id", ParticipantController.DeleteParticipant)

	// beneficiaries
	api.GET("/beneficiaries", BeneficiaryController.GetBeneficiaries)
	api.GET("/beneficiaries/:id", BeneficiaryController.GetBeneficiary)
	api.POST("/beneficiaries", BeneficiaryController.CreateBeneficiary)
	api.PATCH("/beneficiaries/:id", BeneficiaryController.UpdateBeneficiary)
	api.DELETE("/beneficiaries/:id", BeneficiaryController.DeleteBeneficiary)

	// transactions
	api.GET("/transactions", TransactionController.GetTransactions)
	api.GET("/transactions/mosque", TransactionController.GetTransactionsByMosqueID)
	api.POST("/transactions", TransactionController.CreateTransaction)
	api.PUT("/transactions/:id/proof", TransactionController.UploadProof)
	api.PUT("/transactions/:id/verify", TransactionController.VerifyTransaction)
}
