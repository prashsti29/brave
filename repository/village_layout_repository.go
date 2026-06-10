package repository

import (
	"gorm.io/gorm"
	"github.com/prashsti29/brave/models"
)

type VillageLayoutRepository struct {
	database *gorm.DB
}

func NewVillageLayoutRepository(connection *gorm.DB) *VillageLayoutRepository {
	var villageRepo VillageLayoutRepository
	villageRepo.database = connection
	var result *VillageLayoutRepository
	result = &villageRepo
	return result
}

func (villageRepo *VillageLayoutRepository) CreateVillageLayout(layout *models.VillageLayout) error {
	var result *gorm.DB
	result = villageRepo.database.Create(layout)
	return result.Error
}

func (villageRepo *VillageLayoutRepository) GetVillageByPlayerID(playerID string) ([]models.VillageLayout, error) {
	 var layouts []models.VillageLayout
    var result *gorm.DB
    result = villageRepo.database.Where("player_id = ?", playerID).Find(&layouts)
    return layouts, result.Error
}
