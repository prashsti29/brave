package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prashsti29/brave/internal/models"
	"github.com/prashsti29/brave/internal/service"
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

// CreateTroops validates and creates troops for the player
func (buildingController *BuildingController) CreateTroops(responseWriter http.ResponseWriter, request *http.Request) {
	var troopReq struct {
		Name   string `json:"name"`
		Level  int    `json:"level"`
	}
	err := json.NewDecoder(request.Body).Decode(&troopReq)
	if err != nil {
		http.Error(responseWriter, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Create troop config from request
	troopConfig := models.TroopConfig{
		Name:              troopReq.Name,
		Level:             troopReq.Level,
		UnlocksAtDunbrochLevel: 1,
		CostWisps:         10,
		CostEmbis:         0,
		HousingSpace:      1,
		MaxAllowed:        1,
	}

	// Validate troop creation using building service
	if !buildingController.buildingService.ValidateTroopCreation(nil, troopConfig) {
		http.Error(responseWriter, "Invalid troop configuration: insufficient Dunbroch level, gems, or housing space", http.StatusBadRequest)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(map[string]string{"status": "troops created"})
}

// AddBuilding handles creation and placement of a new building
func (buildingController *BuildingController) AddBuilding(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	playerID := vars["player_id"]

	var req struct {
		Type      string `json:"type"`
		Name      string `json:"name"`
		MaxHealth int    `json:"max_health"`
		X         int    `json:"x"`
		Y         int    `json:"y"`
	}

	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		http.Error(responseWriter, "Invalid request body", http.StatusBadRequest)
		return
	}

	building := models.Building{
		PlayerID:      playerID,
		Type:          req.Type,
		Name:          req.Name,
		Level:         1,
		MaxHealth:     req.MaxHealth,
		CurrentHealth: req.MaxHealth,
		DunbrochLevel: 1,
		MaxAllowed:    1,
		IsUpgrading:   false,
	}

	err = buildingController.buildingService.AddBuilding(&building, req.X, req.Y)
	if err != nil {
		http.Error(responseWriter, "Could not add building", http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(responseWriter).Encode(building)
}

// MoveBuilding handles moving an existing building in the village layout
func (buildingController *BuildingController) MoveBuilding(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	playerID := vars["player_id"]
	buildingID := vars["building_id"]

	var req struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		http.Error(responseWriter, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = buildingController.buildingService.MoveBuilding(playerID, buildingID, req.X, req.Y)
	if err != nil {
		http.Error(responseWriter, "Could not move building", http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(map[string]string{"status": "building moved"})
}

// CompleteUpgrade handles completing an upgrade instantly by spending gems
func (buildingController *BuildingController) CompleteUpgrade(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	buildingID := vars["building_id"]

	building, err := buildingController.buildingService.CompleteUpgrade(buildingID)
	if err != nil {
		http.Error(responseWriter, "Could not complete upgrade", http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(building)
}

// UpgradeBuilding handles starting a building upgrade
func (buildingController *BuildingController) UpgradeBuilding(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	buildingID := vars["building_id"]

	building, err := buildingController.buildingService.UpgradeBuilding(buildingID)
	if err != nil {
		http.Error(responseWriter, "Could not upgrade building", http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(building)
}

// CollectGold handles collecting gold from producer buildings
func (buildingController *BuildingController) CollectGold(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	playerID := vars["player_id"]

	gold, err := buildingController.buildingService.CollectGold(playerID)
	if err != nil {
		http.Error(responseWriter, "Could not collect gold", http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(map[string]int{"gold": gold})
}

// CollectElixir handles collecting elixir from producer buildings
func (buildingController *BuildingController) CollectElixir(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	playerID := vars["player_id"]

	elixir, err := buildingController.buildingService.CollectElixir(playerID)
	if err != nil {
		http.Error(responseWriter, "Could not collect elixir", http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(map[string]int{"elixir": elixir})
}

