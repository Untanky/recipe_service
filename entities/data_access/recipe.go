package data_access

import "github.com/Untanky/recipe_service/entities"

type RecipeDataAccess interface {
	GetByID(id int) (*entities.Recipe, error)
	Create(recipe *entities.Recipe) error
	Update(recipe *entities.Recipe) error
	Delete(recipe *entities.Recipe) error
	DeleteByID(id int) error
}
