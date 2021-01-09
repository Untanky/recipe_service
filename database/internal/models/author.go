package models

import (
	"github.com/Untanky/recipe_service/entities"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID             int
	Firstname      string
	Lastname       string
	Location       string
	Description    string
	ProfileImageID string
	Recipes        []Recipe
}

func NewAuthorFromEntity(entity *entities.Author) *Author {
	return &Author{
		ID:             entity.ID,
		Firstname:      entity.Firstname,
		Lastname:       entity.Lastname,
		Description:    entity.Description,
		Location:       entity.Location,
		ProfileImageID: entity.ProfileImageID,
	}
}

func (author *Author) UpdateFromEntity(entity *entities.Author) {
	author.Firstname = entity.Firstname
	author.Lastname = entity.Lastname
	author.Description = entity.Description
	author.Location = entity.Location
	author.ProfileImageID = entity.ProfileImageID
}

func (author *Author) ToEntity() *entities.Author {
	return &entities.Author{
		ID:             author.ID,
		Firstname:      author.Firstname,
		Lastname:       author.Lastname,
		Description:    author.Description,
		Location:       author.Location,
		ProfileImageID: author.ProfileImageID,
	}
}
