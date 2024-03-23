package controllers

import (
	//_ "controller"

	"fmt"
	"log"
	"net/http"

	"github.com/gkganesh126/recipe-sharing-platform/common"
	db "github.com/gkganesh126/recipe-sharing-platform/db-ops"
	"github.com/gkganesh126/recipe-sharing-platform/models/frontend"
	"github.com/gkganesh126/recipe-sharing-platform/sessions"
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
	// Create response db

	err = sessions.SetSession(user.Data.Username, &response)
	if err != nil {
		common.DisplayAppError(response, user.Data.UserID.String(), err, "CreateUser SetSession failed", http.StatusInternalServerError)
		return
	}

	zap.S().Infof("username: %s successfully created", user.Data.Username)
	http.Redirect(response, request, "/internalnew", 302)
}

// login handler

func LoginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")
	log.Println("at LoginHandler: ", name, pass)
	redirectTarget := "/"
	_, err := sessions.IsValidSession(request)
	log.Println("IsValidSession error:", err)
	if err == nil {
		// .. check credentials ..
		sessions.SetSession(name, &response)
		redirectTarget = "/internal"
		http.Redirect(response, request, redirectTarget, 302)
	} else {
		redirectTarget = "/errorlogin"
		http.Redirect(response, request, redirectTarget, 302)
	}
}

// logout handler

func LogoutHandler(response http.ResponseWriter, request *http.Request) {
	sessions.ClearSession(response)
	http.Redirect(response, request, "/", 302)
}

func IndexPageHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "%s %s", frontend.IndexPage, "Login")
}
func IndexPageHandler1(response http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(response, "%s %s", frontend.IndexPage, "Login Failed, Try again")
	if err != nil {
		fmt.Println(err.Error())
	}
}

// internal page

func InternalPageHandler(response http.ResponseWriter, request *http.Request) {
	username, err := sessions.IsValidSession(request)
	log.Println("at InternalPageHandler1", username, err)
	if err == nil {
		fmt.Fprintf(response, "%s", frontend.InternalPage)
	} else {
		http.Redirect(response, request, "/", http.StatusFound)
	}
}
func InternalPageHandler1(response http.ResponseWriter, request *http.Request) {
	username, err := sessions.IsValidSession(request)
	log.Println("at InternalPageHandler1", username, err)
	if err == nil {
		fmt.Fprintf(response, "%s", frontend.InternalPage)
	} else {
		http.Redirect(response, request, "/", http.StatusFound)
	}
}
