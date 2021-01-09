package database

import (
	"github.com/Untanky/recipe_service/database/internal/models"
	"github.com/Untanky/recipe_service/entities"
	"gorm.io/gorm"
)

type LatestSegment struct {
	DB gorm.DB
}

func (segment *LatestSegment) GetRecipes() ([]*entities.RecipePreview, error) {
	const order = "created_at desc"

	offset := 0
	limit := 10

	previewModels := []*models.RecipePreview{}

	result := segment.DB.
		Offset(offset).
		Limit(limit).
		Order(order).
		Find(&previewModels)
	if result.Error != nil {
		return nil, result.Error
	}

	previews := []*entities.RecipePreview{}
	for _, previewModel := range previewModels {
		entity := previewModel.ToEntity()
		previews = append(previews, entity)
	}

	return previews, nil
}
