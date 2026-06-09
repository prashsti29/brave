package repository

import (
	"gorm.io/gorm"
	"github.com/prashsti29/brave/models"
)

type BuildingRepository struct {
	database *gorm.DB
}

func NewBuildingRepository(connection *gorm.DB) *BuildingRepository {
	var buildingRepo BuildingRepository
	buildingRepo.database = connection
	var result *BuildingRepository
	result = &buildingRepo
	return result
}

func (buildingRepo *BuildingRepository) CreateBuilding(building *models.Building) error {
 	var result *gorm.DB
    result = buildingRepo.database.Create(building)
    return result.Error
}

func (buildingRepo *BuildingRepository) GetBuildingsByPlayerID(playerID string) ([]models.Building, error) {
	var buildings []models.Building
	var result *gorm.DB
	result = buildingRepo.database.Where("player_id = ?", playerID).Find(&buildings)
	return buildings, result.Error
}
