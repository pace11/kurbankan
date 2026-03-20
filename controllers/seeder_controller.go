package controllers

import (
	"fmt"
	"io"
	"kurbankan/config"
	"kurbankan/utils"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type SeederController struct{}

func NewSeederController() *SeederController {
	return &SeederController{}
}

var seedFiles = map[string]string{
	"provinces":  "database/seeds/001_202603001_provinces.sql",
	"regencies":  "database/seeds/001_202603002_regencies.sql",
	"districts":  "database/seeds/001_202603003_districts.sql",
	"villages_1": "database/seeds/001_202603004_villages_1.sql",
	"villages_2": "database/seeds/001_202603005_villages_2.sql",
}

var seedOrder = []string{
	"provinces",
	"regencies",
	"districts",
	"villages_1",
	"villages_2",
}

// GetSeederStatus checks the status of seeded data
// @Summary Check seeder status
// @Description Check which tables have been seeded and data counts
// @Tags Seeder
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response
// @Router /api/seed/status [get]
func (sc *SeederController) GetSeederStatus(c *gin.Context) {
	type TableStatus struct {
		Table   string `json:"table"`
		Count   int64  `json:"count"`
		IsEmpty bool   `json:"is_empty"`
		HasData bool   `json:"has_data"`
	}

	tables := []string{"provinces", "regencies", "districts", "villages"}
	statuses := []TableStatus{}

	for _, table := range tables {
		var count int64
		config.DB.Table(table).Count(&count)

		statuses = append(statuses, TableStatus{
			Table:   table,
			Count:   count,
			IsEmpty: count == 0,
			HasData: count > 0,
		})
	}

	allSeeded := true
	for _, status := range statuses {
		if status.IsEmpty {
			allSeeded = false
			break
		}
	}

	response := gin.H{
		"status":     "success",
		"all_seeded": allSeeded,
		"tables":     statuses,
		"total_count": func() int64 {
			var t int64
			for _, s := range statuses {
				t += s.Count
			}
			return t
		}(),
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Seeder status retrieved",
		"data":    response,
	})
}

// RunAllSeeders runs all seed files in order
// @Summary Run all seeders
// @Description Execute all seed SQL files in the correct order
// @Tags Seeder
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/seed/run [post]
func (sc *SeederController) RunAllSeeders(c *gin.Context) {
	// Security check: optional migration key for production
	// migrationKey := c.GetHeader("X-Migration-Key")
	// expectedKey := os.Getenv("MIGRATION_KEY")

	// if expectedKey != "" && migrationKey != expectedKey {
	// 	utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized: Invalid migration key")
	// 	return
	// }

	results := []gin.H{}
	failed := []string{}

	for _, seed := range seedOrder {
		filePath := seedFiles[seed]
		count, err := executeSeedFile(filePath)

		if err != nil {
			failed = append(failed, seed)
			results = append(results, gin.H{
				"seed":   seed,
				"status": "failed",
				"error":  err.Error(),
				"count":  0,
			})
		} else {
			results = append(results, gin.H{
				"seed":   seed,
				"status": "success",
				"count":  count,
			})
		}
	}

	status := "success"
	message := "All seeders executed successfully"
	if len(failed) > 0 {
		status = "partial_failure"
		message = fmt.Sprintf("Some seeders failed: %s", strings.Join(failed, ", "))
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": message,
		"data": gin.H{
			"results": results,
			"failed":  failed,
			"success": len(seedOrder) - len(failed),
			"total":   len(seedOrder),
		},
	})
}

// RunSpecificSeeder runs a specific seed file
// @Summary Run specific seeder
// @Description Execute a specific seed SQL file
// @Tags Seeder
// @Accept json
// @Produce json
// @Param seed path string true "Seed name (provinces, regencies, districts, villages_1, villages_2)"
// @Security ApiKeyAuth
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/seed/run/{seed} [post]
func (sc *SeederController) RunSpecificSeeder(c *gin.Context) {
	seedName := c.Param("seed")

	// Security check
	// migrationKey := c.GetHeader("X-Migration-Key")
	// expectedKey := os.Getenv("MIGRATION_KEY")

	// if expectedKey != "" && migrationKey != expectedKey {
	// 	utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized: Invalid migration key")
	// 	return
	// }

	filePath, exists := seedFiles[seedName]
	if !exists {
		utils.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Unknown seed: %s. Available: %v", seedName, seedOrder))
		return
	}

	count, err := executeSeedFile(filePath)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Seeder failed: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Seeder '%s' executed successfully", seedName),
		"data": gin.H{
			"seed":  seedName,
			"count": count,
		},
	})
}

// ClearSeededData clears all seeded data (USE WITH CAUTION!)
// @Summary Clear all seeded data
// @Description Delete all data from provinces, regencies, districts, and villages tables
// @Tags Seeder
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/seed/clear [delete]
func (sc *SeederController) ClearSeededData(c *gin.Context) {
	// Security check with confirmation
	// migrationKey := c.GetHeader("X-Migration-Key")
	// expectedKey := os.Getenv("MIGRATION_KEY")

	// if expectedKey == "" {
	// 	expectedKey = "kurbankan-migration-secret"
	// }

	// if migrationKey != expectedKey {
	// 	utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized: Invalid migration key")
	// 	return
	// }

	var req struct {
		Confirm string `json:"confirm" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Confirmation required")
		return
	}

	if req.Confirm != "CLEAR_SEEDED_DATA" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid confirmation. Must be 'CLEAR_SEEDED_DATA'")
		return
	}

	// Clear in reverse order to respect foreign keys
	tables := []string{"villages", "districts", "regencies", "provinces"}
	cleared := []gin.H{}

	for _, table := range tables {
		var countBefore int64
		config.DB.Table(table).Count(&countBefore)

		result := config.DB.Exec(fmt.Sprintf("DELETE FROM %s", table))
		if result.Error != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Failed to clear %s: %s", table, result.Error.Error()))
			return
		}

		cleared = append(cleared, gin.H{
			"table":        table,
			"rows_deleted": countBefore,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "All seeded data cleared successfully",
		"data": gin.H{
			"cleared": cleared,
		},
	})
}

// executeSeedFile reads and executes a SQL seed file
func executeSeedFile(filePath string) (int, error) {
	// Get absolute path
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return 0, fmt.Errorf("failed to resolve path: %w", err)
	}

	// Read file
	file, err := os.Open(absPath)
	if err != nil {
		return 0, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}

	sqlContent := string(content)

	// Convert MySQL syntax to PostgreSQL
	// Remove backticks
	sqlContent = strings.ReplaceAll(sqlContent, "`", "")

	// Split by semicolon for multiple statements
	statements := strings.Split(sqlContent, ";")

	insertCount := 0
	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" || strings.HasPrefix(stmt, "--") {
			continue
		}

		// Execute statement
		result := config.DB.Exec(stmt)
		if result.Error != nil {
			return insertCount, fmt.Errorf("failed to execute statement: %w", result.Error)
		}

		if result.RowsAffected > 0 {
			insertCount += int(result.RowsAffected)
		}
	}

	return insertCount, nil
}
