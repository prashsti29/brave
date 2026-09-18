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
	router.HandleFunc("/buildings/{player_id}/{building_id}/move", buildingController.MoveBuilding).Methods("PUT")
	router.HandleFunc("/collect/gold", buildingController.CollectGold).Methods("POST")
	router.HandleFunc("/collect/elixir", buildingController.CollectElixir).Methods("POST")
}
