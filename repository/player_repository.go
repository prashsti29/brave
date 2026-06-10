package repository

import (
	"gorm.io/gorm"
	"github.com/prashsti29/brave/models"
)

type PlayerRepository struct {
	database *gorm.DB
}

func NewPlayerRepository(connection *gorm.DB) *PlayerRepository {
	var playerRepo PlayerRepository
	playerRepo.database = connection
	var result *PlayerRepository
	result = &playerRepo
	return result
}

func (playerRepo *PlayerRepository) CreatePlayer(player *models.Player) error {
	 var result *gorm.DB
    result = playerRepo.database.Create(player)
    return result.Error
}

func (playerRepo *PlayerRepository) GetPlayerByID(id string) (*models.Player, error) {
	var player models.Player
    var result *gorm.DB
    result = playerRepo.database.First(&player, "id = ?", id)
    var playerResult *models.Player
    playerResult = &player
    return playerResult, result.Error
}
