package service

import (
    "github.com/google/uuid"
    "github.com/prashsti29/brave/models"
    "github.com/prashsti29/brave/repository"
)

type BuildingService struct {
    buildingRepo *repository.BuildingRepository
    villageService *VillageLayoutService
}

func NewBuildingService(buildingRepo *repository.BuildingRepository, villageService *VillageLayoutService) *BuildingService {
    var buildingService BuildingService
    buildingService.buildingRepo = buildingRepo
    buildingService.villageService = villageService
    var result *BuildingService
    result = &buildingService
    return result
}

func (buildingService *BuildingService) CreateDefaultBuildings(playerID string) error {
    var err error

    type defaultBuilding struct {
        buildingType string
        name         string
        x            int
        y            int
        maxHealth    int
    }

    var defaults []defaultBuilding
    defaults = []defaultBuilding{
        {"dunbroch", "Dunbroch", 5, 5, 1000},
        {"builders_hut", "Builders Hut", 7, 5, 500},
        {"gold_storage", "Gold Storage", 5, 7, 600},
        {"elixir_storage", "Elixir Storage", 7, 7, 600},
    }

    for _, defaultB := range defaults {
        var building models.Building
        building.ID = uuid.New().String()
        building.PlayerID = playerID
        building.Type = defaultB.buildingType
        building.Name = defaultB.name
        building.Level = 1
        building.MaxHealth = defaultB.maxHealth
        building.CurrentHealth = defaultB.maxHealth
        building.DunbrochLevel = 1
        building.MaxAllowed = 1
        building.IsUpgrading = false

        err = buildingService.buildingRepo.CreateBuilding(&building)
        if err != nil {
            return err
        }

        err = buildingService.villageService.PlaceBuilding(playerID, building.ID, defaultB.x, defaultB.y)
        if err != nil {
            return err
        }
    }

    return nil
}

func (buildingService *BuildingService) GetBuildingsByPlayerID(playerID string) ([]models.Building, error) {
    var buildings []models.Building
    var err error
    buildings, err = buildingService.buildingRepo.GetBuildingsByPlayerID(playerID)
    return buildings, err
}