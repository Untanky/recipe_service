package database

import (
	"github.com/Untanky/recipe_service/database/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(models.Recipe{})
	db.AutoMigrate(models.Author{})
}
