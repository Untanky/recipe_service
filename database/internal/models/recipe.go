package models

import (
	"github.com/Untanky/recipe_service/entities"
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	RecipePreview
	AuthorID    int
	Ingredients []string
	Steps       []string
}

func NewRecipeFromEntity(entity *entities.Recipe) *Recipe {
	return &Recipe{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		AuthorID:    entity.Author.ID,
		Ingredients: entity.Ingredients,
		Steps:       entity.Steps,
	}
}

func (recipe *Recipe) UpdateFromEntity(entity *entities.Recipe) {
	recipe.Title = entity.Title
	recipe.Description = entity.Description
	recipe.AuthorID = entity.Author.ID
	recipe.Ingredients = entity.Ingredients
	recipe.Steps = entity.Steps
}

func (recipe *Recipe) ToEntity() *entities.Recipe {
	return &entities.Recipe{
		ID:          recipe.ID,
		Title:       recipe.Title,
		Description: recipe.Description,
		Ingredients: recipe.Ingredients,
		Steps:       recipe.Steps,
	}
}
