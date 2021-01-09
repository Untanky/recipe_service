package entities

import "github.com/Untanky/recipe_service/entities/data_access"

type Author struct {
	ID             int
	Firstname      string
	Lastname       string
	Username       string
	Location       string
	Description    string
	ProfileImageID string
}

func (author *Author) GetRecipes() (data_access.Segment, error) {
	return nil, nil
}
