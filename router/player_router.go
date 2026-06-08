package router

import (
	"github.com/gorilla/mux"
	"github.com/prashsti29/brave/controllers"
)

func RegisterPlayerRoutes(router *mux.Router, playerController *controllers.PlayerController) {
	router.HandleFunc("/players", playerController.CreatePlayer).Methods("POST")
	router.HandleFunc("/players/{id}", playerController.GetPlayer).Methods("GET")
}
