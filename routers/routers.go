package routers

import (
	"fmt"

	"github.com/gkganesh126/recipe-sharing-platform/controllers"

	//"log"
	"net/http"

	"github.com/gkganesh126/recipe-sharing-platform/models/frontend"

	"github.com/gorilla/mux"
)

func SetRecipeSharingRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controllers.IndexPageHandler)
	router.HandleFunc("/errorlogin", controllers.IndexPageHandler1)
	router.HandleFunc("/internal", controllers.InternalPageHandler)
	router.HandleFunc("/internalnew", controllers.InternalPageHandler1)
	router.HandleFunc("/register/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "%s", frontend.RegisterPage)
	})
	router.HandleFunc("/newlogin", controllers.RegisterInDb)
	router.HandleFunc("/login", controllers.LoginHandler)
	router.HandleFunc("/logout", controllers.LogoutHandler)

	/*router.HandleFunc("/upload/", controllers.HelloServer)

	router.HandleFunc("/app/", controllers.Uploadimage)

	router.HandleFunc("/images/", controllers.HandleImages)
	//r.HandleFunc("/images/{ImageName}/", HandleImage)

	router.HandleFunc("/viewimage/", controllers.Viewimage)
	router.HandleFunc("/writecmnttodb/", controllers.WriteCmntToDb)
	router.HandleFunc("/readcmntfromdb/", controllers.ReadCmntFromDb)
	*/

	return router
}
