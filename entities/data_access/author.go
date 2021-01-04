package data_access

import "github.com/Untanky/recipe_service/entities"

type AuthorDataAccess interface {
	GetByID(id int) (*entities.Author, error)
	Create(author entities.Author) error
	Update(author entities.Author) error
	Delete(author entities.Author) error
	DeleteByID(id int) error
}
