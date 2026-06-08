package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prashsti29/brave/service"
)

type VillageLayoutController struct {
	villageService *service.VillageLayoutService
}

func NewVillageLayoutController(villageService *service.VillageLayoutService) *VillageLayoutController {
	var villageController VillageLayoutController
	villageController.villageService = villageService
	var result *VillageLayoutController
	result = &villageController
	return result
}

func (villageController *VillageLayoutController) GetVillageByPlayerID(responseWriter http.ResponseWriter, request *http.Request) {
	var vars map[string]string
	vars = mux.Vars(request)
	var playerID string
	playerID = vars["player_id"]

	var layouts, err = villageController.villageService.GetVillageByPlayerID(playerID)
	if err != nil {
		http.Error(responseWriter, "Could not fetch village", http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(layouts)
}
