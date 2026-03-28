package controllers

import (
	"kurbankan/config"
	"kurbankan/models"
	"kurbankan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MigrationController struct{}

func NewMigrationController() *MigrationController {
	return &MigrationController{}
}

// RunMigration runs auto-migration for all models
// @Summary Run database migration
// @Description Automatically migrate all database tables based on GORM models
// @Tags Migration
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/migrate [post]
func (mc *MigrationController) RunMigration(c *gin.Context) {
	// Security check: require migration key
	// migrationKey := c.GetHeader("X-Migration-Key")
	// expectedKey := os.Getenv("MIGRATION_KEY")

	// // If MIGRATION_KEY is not set in .env, use a default for development
	// if expectedKey == "" {
	// 	expectedKey = "kurbankan-migration-secret"
	// }

	// if migrationKey != expectedKey {
	// 	utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized: Invalid migration key")
	// 	return
	// }

	// List of all models to migrate
	modelsToMigrate := []interface{}{
		// Area/Wilayah models
		&models.Province{},
		&models.Regency{},
		&models.District{},
		&models.Village{},

		// User & Auth models
		&models.User{},

		// Mosque models
		&models.Mosque{},
		&models.MosqueMember{},

		// Participant models
		&models.Participant{},

		// Qurban Period & Offerings
		&models.QurbanPeriod{},
		&models.QurbanOffering{},

		// Qurban Animals
		&models.QurbanAnimal{},
		&models.AnimalParticipant{},
		&models.AnimalMedia{},
		&models.AnimalLog{},

		// Transaction models
		&models.Transaction{},
		&models.TransactionItem{},

		// Beneficiary & Distribution models
		&models.Beneficiary{},
		&models.DistributionBatch{},
		&models.DistributionItem{},

		// Reports
		&models.Report{},
	}

	// Run auto-migration
	for _, model := range modelsToMigrate {
		if err := config.DB.AutoMigrate(model); err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Migration failed: "+err.Error())
			return
		}
	}

	// Get table names for response
	tableNames := []string{
		"provinces",
		"regencies",
		"districts",
		"villages",
		"users",
		"mosques",
		"mosque_members",
		"participants",
		"qurban_periods",
		"qurban_options",
		"qurban_animals",
		"animal_participants",
		"animal_media",
		"animal_logs",
		"transactions",
		"transaction_items",
		"beneficiaries",
		"distribution_batches",
		"distribution_items",
		"reports",
	}

	response := gin.H{
		"message": "Database migration completed successfully",
		"tables":  tableNames,
		"count":   len(tableNames),
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Migration successful",
		"data":    response,
	})
}

// CheckMigrationStatus checks if all tables exist
// @Summary Check migration status
// @Description Check if all required database tables exist
// @Tags Migration
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response
// @Router /api/migrate/status [get]
func (mc *MigrationController) CheckMigrationStatus(c *gin.Context) {
	tableNames := []string{
		"provinces",
		"regencies",
		"districts",
		"villages",
		"users",
		"mosques",
		"mosque_members",
		"participants",
		"qurban_periods",
		"qurban_options",
		"qurban_animals",
		"animal_participants",
		"animal_media",
		"animal_logs",
		"transactions",
		"transaction_items",
		"beneficiaries",
		"distribution_batches",
		"distribution_items",
		"reports",
	}

	existingTables := []string{}
	missingTables := []string{}

	for _, tableName := range tableNames {
		if config.DB.Migrator().HasTable(tableName) {
			existingTables = append(existingTables, tableName)
		} else {
			missingTables = append(missingTables, tableName)
		}
	}

	status := "complete"
	if len(missingTables) > 0 {
		status = "incomplete"
	}

	response := gin.H{
		"status":          status,
		"total_tables":    len(tableNames),
		"existing_tables": existingTables,
		"missing_tables":  missingTables,
		"need_migration":  len(missingTables) > 0,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Migration status retrieved",
		"data":    response,
	})
}

// DropAllTables drops all tables (USE WITH CAUTION!)
// @Summary Drop all database tables
// @Description DROP ALL TABLES - USE WITH EXTREME CAUTION! This will delete all data.
// @Tags Migration
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/migrate/drop [delete]
func (mc *MigrationController) DropAllTables(c *gin.Context) {
	// Security check: require migration key AND confirmation
	// migrationKey := c.GetHeader("X-Migration-Key")
	// expectedKey := os.Getenv("MIGRATION_KEY")

	// if expectedKey == "" {
	// 	expectedKey = "kurbankan-migration-secret"
	// }

	// if migrationKey != expectedKey {
	// 	utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized: Invalid migration key")
	// 	return
	// }

	// Additional confirmation required
	var req struct {
		Confirm string `json:"confirm" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Confirmation required")
		return
	}

	if req.Confirm != "DROP_ALL_TABLES" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid confirmation. Must be 'DROP_ALL_TABLES'")
		return
	}

	// Drop tables in reverse order to handle foreign key constraints
	modelsToMigrate := []interface{}{
		// Drop tables with the most dependencies first
		&models.Report{},
		&models.DistributionItem{},
		&models.DistributionBatch{},
		&models.AnimalLog{},
		&models.AnimalMedia{},
		&models.AnimalParticipant{},
		&models.QurbanAnimal{},
		&models.Beneficiary{},
		&models.TransactionItem{},
		&models.Transaction{},
		&models.QurbanOffering{},
		&models.QurbanPeriod{},
		&models.MosqueMember{},
		&models.Participant{},
		&models.Mosque{},
		&models.User{},
		&models.Village{},
		&models.District{},
		&models.Regency{},
		&models.Province{},
	}

	droppedTables := []string{}

	for _, model := range modelsToMigrate {
		if err := config.DB.Migrator().DropTable(model); err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to drop table: "+err.Error())
			return
		}
		// Get table name using type assertion
		if m, ok := model.(interface{ TableName() string }); ok {
			droppedTables = append(droppedTables, m.TableName())
		}
	}

	response := gin.H{
		"message":        "All tables dropped successfully",
		"dropped_tables": droppedTables,
		"count":          len(droppedTables),
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Tables dropped",
		"data":    response,
	})
}
