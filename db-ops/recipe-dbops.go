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
