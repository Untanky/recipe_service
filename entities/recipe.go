package entities

type Recipe struct {
	ID          int
	Title       string
	Description string
	Author      Author
	Ingredients []string
	Steps       []string
}
