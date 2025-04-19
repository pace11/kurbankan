package utils

import "gorm.io/gorm"

func FilterByParams(db *gorm.DB, filters map[string]interface{}) *gorm.DB {
	for key, val := range filters {
		if key == "name" {
			if str, ok := val.(string); ok && str != "" {
				db = db.Where("name LIKE ?", "%"+str+"%")
			}
		} else if str, ok := val.(string); ok && str != "" {
			db = db.Where(key+" = ?", str)
		}
	}
	return db
}
