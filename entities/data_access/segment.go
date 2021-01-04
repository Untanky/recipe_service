package data_access

import "github.com/Untanky/recipe_service/entities"

type Segment interface {
	GetRecipes() ([]*entities.Recipe, error)
}

type LatestSegment struct {
}

func (segment *LatestSegment) GetRecipes() ([]*entities.Recipe, error) {
	return nil, nil
}

type PopularSegment struct {
}

func (segment *PopularSegment) GetRecipes() ([]*entities.Recipe, error) {
	return nil, nil
}

type RecommendedSegment struct {
}

func (segment *RecommendedSegment) GetRecipes() ([]*entities.Recipe, error) {
	return nil, nil
}
