package database

import (
	"github.com/Untanky/recipe_service/entities"
	"gorm.io/gorm"
)

type AuthorDatabase struct {
	DB gorm.DB
}

func (database *AuthorDatabase) GetByID(id int) (*entities.Author, error) {
	// recipeModel := models.Recipe{}

	// result := database.DB.First(&recipeModel, id)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	// recipe := recipeModel.ToEntity()
	// return recipe, nil
	return nil, nil
}

func (database *AuthorDatabase) Create(recipe *entities.Author) error {
	// recipeModel := models.NewRecipeFromEntity(recipe)

	// database.DB.Create(recipeModel)
	// return database.DB.Error
	return nil
}

func (database *AuthorDatabase) Update(recipe *entities.Author) error {
	// recipeModel := models.Recipe{}

	// database.DB.First(&recipeModel, recipe.ID)
	// if database.DB.Error != nil {
	// 	return database.DB.Error
	// }

	// recipeModel.UpdateFromEntity(recipe)

	// database.DB.Save(recipeModel)
	// return database.DB.Error
	return nil
}

func (database *AuthorDatabase) Delete(recipe *entities.Author) error {
	return database.DeleteByID(recipe.ID)
}

func (database *AuthorDatabase) DeleteByID(id int) error {
	// database.DB.Delete(&models.Recipe{}, id)
	// return database.DB.Error
	return nil
}
