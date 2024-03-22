package main

import (
	"log"
	"net/http"

	"github.com/gkganesh126/recipe-sharing-platform/controllers"
	"github.com/gkganesh126/recipe-sharing-platform/routers"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(false)
	// Routes for the User entity
	router = routers.SetRecipeSharingRouters(router)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", router)

	controllers.InitCookieHandler()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	//return router
}
