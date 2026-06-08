package controllers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/prashsti29/brave/service"
)

type BuildingController struct {
    buildingService *service.BuildingService
}

func NewBuildingController(buildingService *service.BuildingService) *BuildingController {
    var buildingController BuildingController
    buildingController.buildingService = buildingService
    var result *BuildingController
    result = &buildingController
    return result
}

func (buildingController *BuildingController) GetBuildingsByPlayerID(responseWriter http.ResponseWriter, request *http.Request) {
    var vars map[string]string
    vars = mux.Vars(request)
    var playerID string
    playerID = vars["player_id"]

    var buildings, err = buildingController.buildingService.GetBuildingsByPlayerID(playerID)
    if err != nil {
        http.Error(responseWriter, "Could not fetch buildings", http.StatusInternalServerError)
        return
    }

    responseWriter.Header().Set("Content-Type", "application/json")
    json.NewEncoder(responseWriter).Encode(buildings)
}