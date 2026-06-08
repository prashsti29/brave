package router

import (
	"github.com/gorilla/mux"
	"github.com/prashsti29/brave/controllers"
)

func RegisterBuildingRoutes(router *mux.Router, buildingController *controllers.BuildingController) {
	router.HandleFunc("/buildings/{player_id}", buildingController.GetBuildingsByPlayerID).Methods("GET")
}
