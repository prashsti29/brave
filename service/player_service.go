package service

import (
	"github.com/google/uuid"
	"github.com/prashsti29/brave/models"
	"github.com/prashsti29/brave/repository"
)

type PlayerService struct {
	playerRepo      *repository.PlayerRepository
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
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	player := models.Player{
		ID:           uuid.New().String(),
		Email:        email,
		PasswordHash: hashedPassword,
		DunbrochLevel: 1,
		Gems:         10,
		Wisps:        500,
		Embis:        500,
	}

	err = playerService.playerRepo.CreatePlayer(&player)
	if err != nil {
		return nil, err
	}

	err = playerService.buildingService.CreateDefaultBuildings(player.ID)
	if err != nil {
		return nil, err
	}

	return &player, nil
}

func (playerService *PlayerService) GetPlayerByID(id string) (*models.Player, error) {
	var player *models.Player
	var err error
	player, err = playerService.playerRepo.GetPlayerByID(id)
	return player, err
}

func (playerService *PlayerService) LoginPlayer(email, password string) (string, error) {
	player, err := playerService.playerRepo.GetPlayerByEmail(email)
	if err != nil {
		return "", err
	}

	err = comparePassword(password, player.PasswordHash)
	if err != nil {
		return "", err
	}

	token, err := GenerateToken(player.ID, player.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (playerService *PlayerService) DeletePlayer(id string) error {
	err := playerService.playerRepo.DeletePlayer(id)
	if err != nil {
		return err
	}

	return nil
}
