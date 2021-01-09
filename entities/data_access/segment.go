package data_access

import "github.com/Untanky/recipe_service/entities"

type Segment interface {
	GetRecipes() ([]*entities.RecipePreview, error)
}
