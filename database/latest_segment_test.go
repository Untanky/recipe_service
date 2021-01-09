package database_test

import (
	"regexp"
	"testing"

	"github.com/Untanky/recipe_service/database"
	"github.com/Untanky/recipe_service/database/internal/models"
)

func TestLatestSegmentGetRecipe(t *testing.T) {
	db, mock := SetupDB(t)
	defer CleanUpDB(t)

	segment := database.LatestSegment{
		DB: *db,
	}

	db.AutoMigrate(models.RecipePreview{})

	preview := models.RecipePreview{
		ID:          0,
		Title:       "Test Recipe",
		Description: "Recipe description, that might be longer",
	}

	rows := mock.
		NewRows([]string{"id", "title", "description"}).
		AddRow(preview.ID, preview.Title, preview.Description)

	const query = `SELECT * FROM "recipes" ORDER BY created_at desc LIMIT 10`

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	dbPreviews, err := segment.GetRecipes()
	if err != nil {
		t.Error(err)
		return
	}

	if len(dbPreviews) != 1 {
		t.Error("previews not one element long")
		return
	}

	if *dbPreviews[0] != *preview.ToEntity() {
		t.Errorf("does not match inserted model; expected: %v, actual: %v", dbPreviews[0], preview.ToEntity())
		return
	}
}
