package main

import (
	"log"

	"github.com/Untanky/recipe_service/database"
	"github.com/Untanky/recipe_service/entities"
	"github.com/Untanky/recipe_service/entities/data_access"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, fetchErr := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if fetchErr != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.Migrate(db)

	database := &database.RecipeDatabase{
		DB: *db,
	}

	test(database)
}

func test(database data_access.RecipeDataAccess) {
	recipe := &entities.Recipe{Title: "Test recipe"}
	createErr := database.Create(recipe)
	if createErr != nil {
		log.Panicln(createErr)
	}
	recipe2 := &entities.Recipe{Title: "Test recipe2"}
	createErr = database.Create(recipe2)
	if createErr != nil {
		log.Panicln(createErr)
	}

	fetchedRecipe, fetchErr := database.GetByID(1)
	if fetchErr != nil {
		log.Panicln(fetchErr)
	}
	log.Println(fetchedRecipe)

	updateRecipe := &entities.Recipe{ID: 1, Title: "Hello World"}
	updateErr := database.Update(updateRecipe)
	if updateErr != nil {
		log.Panicln(updateErr)
	}
	log.Println(fetchedRecipe)

	fetchedRecipe, fetchErr = database.GetByID(1)
	if fetchErr != nil {
		log.Panicln(fetchErr)
	}
	log.Println(fetchedRecipe)

	deleteErr := database.Delete(fetchedRecipe)
	if deleteErr != nil {
		log.Panicln(deleteErr)
	}
	log.Println("deleted correcly")

	fetchedRecipe, fetchErr = database.GetByID(1)
	if fetchErr == nil {
		log.Println(fetchedRecipe)
		log.Panicln("should have had error")
	}
	log.Println("no recipe found")

	deleteErr = database.DeleteByID(2)
	if deleteErr != nil {
		log.Panicln(deleteErr)
	}
	log.Println("deleted correcly")

	fetchedRecipe, fetchErr = database.GetByID(2)
	if fetchErr == nil {
		log.Panicln("should have had error")
	}
	log.Println("no recipe found")
}
