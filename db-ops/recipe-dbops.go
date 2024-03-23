package data

import (
	"github.com/gkganesh126/recipe-sharing-platform/models"
	"gopkg.in/mgo.v2"
)

type RecipeRepository struct {
	C *mgo.Collection
}

func (r *RecipeRepository) Create(recipe *models.Recipe) error {
	_, err := r.C.UpsertId(recipe.RecipeID, recipe)
	//err := r.C.Insert(&user)
	return err
}
func (r *RecipeRepository) GetAll() []models.Recipe {
	var recipes []models.Recipe
	iter := r.C.Find(nil).Iter()
	result := models.Recipe{}
	for iter.Next(&result) {
		recipes = append(recipes, result)
	}
	return recipes
}
