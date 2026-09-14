package router

import (
	"github.com/gorilla/mux"
	"github.com/prashsti29/brave/internal/controllers"
)

func RegisterBuildingRoutes(router *mux.Router, buildingController *controllers.BuildingController) {
	router.HandleFunc("/buildings/{player_id}", buildingController.GetBuildingsByPlayerID).Methods("GET")
	router.HandleFunc("/buildings/{player_id}", buildingController.AddBuilding).Methods("POST")
	router.HandleFunc("/buildings/{player_id}/troops", buildingController.CreateTroops).Methods("POST")
	router.HandleFunc("/buildings/{building_id}/upgrade", buildingController.UpgradeBuilding).Methods("PUT")
	router.HandleFunc("/buildings/{building_id}/complete", buildingController.CompleteUpgrade).Methods("POST")
}
