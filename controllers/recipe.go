package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/gkganesh126/recipe-sharing-platform/common"
	db "github.com/gkganesh126/recipe-sharing-platform/db-ops"
	"github.com/gkganesh126/recipe-sharing-platform/models"
	"github.com/gkganesh126/recipe-sharing-platform/models/frontend"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
)

func UploadRecipe(response http.ResponseWriter, request *http.Request) {
	zap.S().Info("At UploadRecipe")

	idTextD := request.FormValue("idtextd")
	rname := request.FormValue("rname")
	file, handler, err := request.FormFile("id-file-d")
	if err != nil {
		zap.S().Error(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		zap.S().Error(err)
	}
	currenttime := time.Now().Nanosecond()
	fname := strconv.Itoa(int(currenttime)) + "_" + handler.Filename
	zap.S().Info("current time is ", currenttime)
	zap.S().Info("fname is ", fname)

	serverPicture := "static/server/pictures/" + fname
	err = os.WriteFile(serverPicture, data, 0777)
	if err != nil {
		zap.S().Error(err)
	}

	var recipe models.Recipe
	recipe.RecipeID = bson.NewObjectId()
	recipe.RecipeName = rname
	recipe.RecipeDetail = idTextD
	recipe.ImageName = fname
	recipe.CurrentUsername = "User"
	recipe.Comments = nil
	zap.S().Info(recipe)

	context := NewContext()
	defer context.Close()
	c := context.RecipeSharingPlatformDbCollection("recipes")
	// Create User
	repo := &db.RecipeRepository{C: c}

	err = repo.Create(&recipe)
	if err != nil {
		common.DisplayAppError(response, "User", err, "CreateRecipe write db failed", http.StatusInternalServerError)
		return
	}

}

func ViewRecipe(response http.ResponseWriter, request *http.Request) {
	zap.S().Info("At ViewRecipe")

	imagee := models.Imagee{}
	t := template.New("Imagee template")
	t, err := t.Parse(frontend.ViewRecipe)
	if err != nil {
		common.DisplayAppError(response, "User", err, "ViewRecipe failed", http.StatusInternalServerError)
		return
	}
	context := NewContext()
	defer context.Close()
	c := context.RecipeSharingPlatformDbCollection("recipes")
	// Create User
	repo := &db.RecipeRepository{C: c}

	vs := repo.GetAll()
	//zap.S().Info("vs: ", vs)
	for i, v := range vs {
		imagee.RecipeID = append(imagee.RecipeID, v.RecipeID)
		imagee.RecipeName = append(imagee.RecipeName, v.RecipeName)
		imagee.RecipeDetail = append(imagee.RecipeDetail, v.RecipeDetail)
		imagee.ImageName = append(imagee.ImageName, v.ImageName)
		for _, comment := range v.Comments {
			imagee.Comments[i] = append(imagee.Comments[i], comment)
		}
	}
	err = t.Execute(response, imagee)
	if err != nil {
		common.DisplayAppError(response, "User", err, "ViewRecipe failed", http.StatusInternalServerError)
		return
	}

}

func WriteCmntToDb(response http.ResponseWriter, request *http.Request) {
	zap.S().Info("At WriteCmntToDb")

	currentImage := request.FormValue("tesla")
	currentComment := request.FormValue("currentComment")

	context := NewContext()
	defer context.Close()
	c := context.RecipeSharingPlatformDbCollection("recipes")
	// Create User
	repo := &db.RecipeRepository{C: c}

	comment := models.Comment{User: "User", Comment: currentComment}
	err := repo.UpdateComments(currentImage, comment)
	if err != nil {
		common.DisplayAppError(response, "User", err, "WriteCmntToDb failed", http.StatusInternalServerError)
		return
	}
}

func ReadCmntFromDb(response http.ResponseWriter, request *http.Request) {
	zap.S().Info("At ReadCmntFromDb")
	currentImage := request.FormValue("tesla")

	context := NewContext()
	defer context.Close()
	c := context.RecipeSharingPlatformDbCollection("recipes")
	// Create User
	repo := &db.RecipeRepository{C: c}
	recipes, err := repo.GetRecipeFromImageName(currentImage)
	if err != nil {
		common.DisplayAppError(response, "User", err, "ReadCmntFromDb failed", http.StatusInternalServerError)
		return
	}
	fmt.Println("recipes: ", recipes)

	data := "<h3>" + recipes.RecipeName + "</h3>" + "<br>" + recipes.RecipeDetail + "<br>"
	for _, comment := range recipes.Comments {
		data += "<h3>" + comment.User + "</h3><br>"
		data += comment.Comment + "<br> <hr>"
	}
	fmt.Println(data)
	fmt.Fprintf(response, data)
}
