package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prashsti29/brave/controllers"
	"github.com/prashsti29/brave/middleware"
)

func RegisterPlayerRoutes(router *mux.Router, playerController *controllers.PlayerController) {
	router.HandleFunc("/auth/register", playerController.Register).Methods("POST")
	router.HandleFunc("/auth/login", playerController.Login).Methods("POST")
	
	router.Handle("/players", middleware.AuthMiddleware(http.HandlerFunc(playerController.CreatePlayer))).Methods("POST")
	router.HandleFunc("/players/{id}", playerController.GetPlayer).Methods("GET")
	
	router.Handle("/profile", middleware.AuthMiddleware(http.HandlerFunc(playerController.GetProfile))).Methods("GET")
}
