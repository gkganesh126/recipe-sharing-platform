package data

import (
	"fmt"

	"github.com/gkganesh126/recipe-sharing-platform/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func (r *RecipeRepository) GetRecipeFromImageName(imageName string) (models.Recipe, error) {
	var recipe models.Recipe
	err := r.C.Find(bson.M{"imagename": imageName}).Select(bson.M{}).One(&recipe)
	if err != nil {
		return models.Recipe{}, err
	}
	return recipe, nil
}
func (r *RecipeRepository) UpdateComments(imageName string, Comment models.Comment) error {
	fmt.Println("imageName: ", imageName, " Comment: ", Comment)
	return r.C.Update(bson.M{"imagename": imageName}, bson.M{"$addToSet": bson.M{"comments": Comment}})
}
