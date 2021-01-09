package models

import "github.com/Untanky/recipe_service/entities"

type RecipePreview struct {
	ID          int
	Title       string
	Description string
}

func (RecipePreview) TableName() string {
	return "recipes"
}

func (recipe *RecipePreview) ToEntity() *entities.RecipePreview {
	return &entities.RecipePreview{
		ID:          recipe.ID,
		Title:       recipe.Title,
		Description: recipe.Description,
	}
}
