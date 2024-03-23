package controllers

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gkganesh126/recipe-sharing-platform/common"
	db "github.com/gkganesh126/recipe-sharing-platform/db-ops"
	"github.com/gkganesh126/recipe-sharing-platform/models"
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
}

func WriteCmntToDb(response http.ResponseWriter, request *http.Request) {
	zap.S().Info("At WriteCmntToDb")
}

func ReadCmntFromDb(response http.ResponseWriter, request *http.Request) {
	zap.S().Info("At ReadCmntFromDb")
}
