package routers

import (
	"fmt"
	"net/http"

	"github.com/gkganesh126/recipe-sharing-platform/controllers"
	"github.com/gkganesh126/recipe-sharing-platform/models/frontend"

	//"log"

	"github.com/gorilla/mux"
)

func SetRecipeSharingRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "%s", frontend.IndexPage)
	})
	router.HandleFunc("/internal", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "%s", frontend.InternalPage)
	})
	router.HandleFunc("/logout", controllers.LogoutHandler)

	router.HandleFunc("/app/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "%s", frontend.AppHtml)
	})
	router.HandleFunc("/upload/", controllers.UploadRecipe)
	router.HandleFunc("/view/", func(response http.ResponseWriter, request *http.Request) {
		controllers.ViewRecipe(response, request)
	})
	router.HandleFunc("/writecmnttodb/", controllers.WriteCmntToDb)
	router.HandleFunc("/readcmntfromdb/", controllers.ReadCmntFromDb)

	return router
}
