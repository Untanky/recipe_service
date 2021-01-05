package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	ID             int
	Firstname      string
	Lastname       string
	Username       string
	Location       string
	Description    string
	ProfileImageID string
	Recipes        []Recipe
}
