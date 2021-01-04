package entities

type Recipe struct {
	ID          int
	Title       string
	Description string
	Author      Author
	Ingredient  []string
	Steps       []string
}
