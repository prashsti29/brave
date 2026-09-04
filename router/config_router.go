package router

import (
	"github.com/gorilla/mux"
	"github.com/prashsti29/brave/controllers"
)

func RegisterConfigRoutes(router *mux.Router, configController *controllers.ConfigController) {
	router.HandleFunc("/configs", configController.GetGameConfigs).Methods("GET")
}