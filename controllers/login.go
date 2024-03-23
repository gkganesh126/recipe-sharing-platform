package controllers

import (
	//_ "controller"

	"log"
	"net/http"

	"github.com/gkganesh126/recipe-sharing-platform/common"
	db "github.com/gkganesh126/recipe-sharing-platform/db-ops"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
)

func RegisterInDb(response http.ResponseWriter, request *http.Request) {

	zap.S().Info("At CreateUser")

	name := request.FormValue("name")
	pass := request.FormValue("password")
	log.Println("at RegisterInDb: ", name, pass)

	context := NewContext()
	defer context.Close()
	c := context.RecipeSharingPlatformDbCollection("users")
	// Create User
	repo := &db.UserRepository{C: c}

	var user UserResource
	user.Data.UserID = bson.NewObjectId()
	user.Data.Username = name
	user.Data.Password = pass
	err := repo.Create(&user.Data)
	if err != nil {
		common.DisplayAppError(response, user.Data.Username, err, "CreateUser write db failed", http.StatusInternalServerError)
		return
	}
	zap.S().Infof("username: %s successfully created", user.Data.Username)
	http.Redirect(response, request, "/internal", http.StatusFound)
}

// logout handler

func LogoutHandler(response http.ResponseWriter, request *http.Request) {
	http.Redirect(response, request, "/", http.StatusFound)
}
