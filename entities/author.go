package entities

type Author struct {
	ID             int
	Firstname      string
	Lastname       string
	Username       string
	Location       string
	Description    string
	ProfileImageID string
}

func (author *Author) GetRecipes() ([]*Recipe, error) {
	return nil, nil
}
