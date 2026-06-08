package router

import (
	"github.com/gorilla/mux"
	"github.com/prashsti29/brave/controllers"
)

func RegisterVillageRoutes(router *mux.Router, villageController *controllers.VillageLayoutController) {
	router.HandleFunc("/village/{player_id}", villageController.GetVillageByPlayerID).Methods("GET")
}
