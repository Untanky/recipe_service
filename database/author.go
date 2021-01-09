package database

import (
	"github.com/Untanky/recipe_service/database/internal/models"
	"github.com/Untanky/recipe_service/entities"
	"gorm.io/gorm"
)

type AuthorDatabase struct {
	DB gorm.DB
}

func (database *AuthorDatabase) GetByID(id int) (*entities.Author, error) {
	authorModel := models.Author{}

	result := database.DB.First(&authorModel, id)
	if result.Error != nil {
		return nil, result.Error
	}

	author := authorModel.ToEntity()
	return author, nil
}

func (database *AuthorDatabase) Create(author *entities.Author) error {
	authorModel := models.NewAuthorFromEntity(author)

	database.DB.Create(authorModel)
	return database.DB.Error
}

func (database *AuthorDatabase) Update(author *entities.Author) error {
	authorModel := models.Author{}

	database.DB.First(&authorModel, author.ID)
	if database.DB.Error != nil {
		return database.DB.Error
	}

	authorModel.UpdateFromEntity(author)

	database.DB.Save(authorModel)
	return database.DB.Error
}

func (database *AuthorDatabase) Delete(author *entities.Author) error {
	return database.DeleteByID(author.ID)
}

func (database *AuthorDatabase) DeleteByID(id int) error {
	database.DB.Delete(&models.Author{}, id)
	return database.DB.Error
	return nil
}
