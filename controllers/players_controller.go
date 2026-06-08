package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prashsti29/brave/service"
)

type CreatePlayerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PlayerController struct {
	playerService *service.PlayerService
}

type PlayerResponse struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	DunbrochLevel int    `json:"dunbroch_level"`
	Gems          int    `json:"gems"`
	Wisps         int    `json:"wisps"`
	Embis         int    `json:"embis"`
}

func NewPlayerController(playerService *service.PlayerService) *PlayerController {
	var playerController PlayerController
	playerController.playerService = playerService
	var result *PlayerController
	result = &playerController
	return result
}

func (playerController *PlayerController) CreatePlayer(responseWriter http.ResponseWriter, request *http.Request) {
	var requestBody CreatePlayerRequest
	var err error

	err = json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		http.Error(responseWriter, "Invalid request body", http.StatusBadRequest)
		return
	}

	player, err := playerController.playerService.CreatePlayer(requestBody.Email, requestBody.Password)
	if err != nil {
		http.Error(responseWriter, "Could not create player", http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusCreated)
	var response PlayerResponse
	response.ID = player.ID
	response.Email = player.Email
	response.DunbrochLevel = player.DunbrochLevel
	response.Gems = player.Gems
	response.Wisps = player.Wisps
	response.Embis = player.Embis
	json.NewEncoder(responseWriter).Encode(response)
}

func (playerController *PlayerController) GetPlayer(responseWriter http.ResponseWriter, request *http.Request) {
	var vars map[string]string
	vars = mux.Vars(request)
	var playerID string
	playerID = vars["id"]

	var player, err = playerController.playerService.GetPlayerByID(playerID)
	if err != nil {
		http.Error(responseWriter, "Player not found", http.StatusNotFound)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(player)
}
