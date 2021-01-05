package database

import (
	"github.com/Untanky/recipe_service/database/models"
	"github.com/Untanky/recipe_service/entities"
	"gorm.io/gorm"
)

type RecipeDatabase struct {
	DB gorm.DB
}

func (database *RecipeDatabase) GetByID(id int) (*entities.Recipe, error) {
	recipeModel := models.Recipe{}

	result := database.DB.First(&recipeModel, id)
	if result.Error != nil {
		return nil, result.Error
	}

	recipe := recipeModel.ToEntity()
	return recipe, nil
}

func (database *RecipeDatabase) Create(recipe *entities.Recipe) error {
	recipeModel := models.NewRecipeFromEntity(recipe)

	database.DB.Create(recipeModel)
	return database.DB.Error
}

func (database *RecipeDatabase) Update(recipe *entities.Recipe) error {
	recipeModel := models.Recipe{}

	database.DB.First(&recipeModel, recipe.ID)
	if database.DB.Error != nil {
		return database.DB.Error
	}

	recipeModel.UpdateFromEntity(recipe)

	database.DB.Save(recipeModel)
	return database.DB.Error
}

func (database *RecipeDatabase) Delete(recipe *entities.Recipe) error {
	return database.DeleteByID(recipe.ID)
}

func (database *RecipeDatabase) DeleteByID(id int) error {
	database.DB.Delete(&models.Recipe{}, id)
	return database.DB.Error
}
