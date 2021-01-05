package models

import (
	"github.com/Untanky/recipe_service/entities"
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	ID    int
	Title string
}

func NewRecipeFromEntity(entity *entities.Recipe) *Recipe {
	return &Recipe{
		ID:    entity.ID,
		Title: entity.Title,
	}
}

func (recipe *Recipe) UpdateFromEntity(entity *entities.Recipe) {
	recipe.Title = entity.Title
}

func (recipe *Recipe) ToEntity() *entities.Recipe {
	return &entities.Recipe{
		ID:    recipe.ID,
		Title: recipe.Title,
	}
}
