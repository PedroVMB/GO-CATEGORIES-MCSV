package database

import (
	c_models "github.com/PedroVMB/GO-CATEGORIES-MCSV/internal/categories/models"
	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(&c_models.Category{})
}
