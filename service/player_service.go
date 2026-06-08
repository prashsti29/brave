package service

import (
    "github.com/google/uuid"
    "github.com/prashsti29/brave/models"
    "github.com/prashsti29/brave/repository"
    "golang.org/x/crypto/bcrypt"
)

type PlayerService struct {
    playerRepo   *repository.PlayerRepository
    buildingService *BuildingService
}

func NewPlayerService(playerRepo *repository.PlayerRepository, buildingService *BuildingService) *PlayerService {
    var playerService PlayerService
    playerService.playerRepo = playerRepo
    playerService.buildingService = buildingService
    var result *PlayerService
    result = &playerService
    return result
}

func (playerService *PlayerService) CreatePlayer(email string, password string) (*models.Player, error) {
    var err error

    var hashedBytes []byte
    hashedBytes, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    var player models.Player
    player.ID = uuid.New().String()
    player.Email = email
    player.PasswordHash = string(hashedBytes)
    player.DunbrochLevel = 1
    player.Gems = 0
    player.Wisps = 0
    player.Embis = 0

    err = playerService.playerRepo.CreatePlayer(&player)
    if err != nil {
        return nil, err
    }

    err = playerService.buildingService.CreateDefaultBuildings(player.ID)
    if err != nil {
        return nil, err
    }

    var result *models.Player
    result = &player
    return result, nil
}

func (playerService *PlayerService) GetPlayerByID(id string) (*models.Player, error) {
    var player *models.Player
    var err error
    player, err = playerService.playerRepo.GetPlayerByID(id)
    return player, err
}