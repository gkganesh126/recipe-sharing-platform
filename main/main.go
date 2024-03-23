package main

import (
	"net/http"

	"github.com/gkganesh126/recipe-sharing-platform/common"
	"github.com/gkganesh126/recipe-sharing-platform/logger"
	"github.com/gkganesh126/recipe-sharing-platform/routers"
	"go.uber.org/zap"

	"github.com/gorilla/mux"
)

func main() {
	loggerMgr := logger.InitLogger()
	zap.ReplaceGlobals(loggerMgr)
	loggerMgr.Sync()

	// Calls startup logic
	common.StartUp()

	router := mux.NewRouter().StrictSlash(false)
	// Routes for the User entity
	router = routers.SetRecipeSharingRouters(router)
	// Get the mux router object

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static"))))
	http.Handle("/", router)

	zap.S().Info("Listening at " + server.Addr)
	server.ListenAndServe()
}
