package service

import (
    "github.com/prashsti29/brave/models"
    "github.com/prashsti29/brave/repository"
)

type VillageLayoutService struct {
    villageRepo *repository.VillageLayoutRepository
}

func NewVillageLayoutService(villageRepo *repository.VillageLayoutRepository) *VillageLayoutService {
    var villageService VillageLayoutService
    villageService.villageRepo = villageRepo
    var result *VillageLayoutService
    result = &villageService
    return result
}

func (villageService *VillageLayoutService) PlaceBuilding(playerID string, buildingID string, x int, y int) error {
    var layout models.VillageLayout
    layout.PlayerID = playerID
    layout.BuildingID = buildingID
    layout.X = x
    layout.Y = y
    var err error
    err = villageService.villageRepo.CreateVillageLayout(&layout)
    return err
}

func (villageService *VillageLayoutService) GetVillageByPlayerID(playerID string) ([]models.VillageLayout, error) {
    var layouts []models.VillageLayout
    var err error
    layouts, err = villageService.villageRepo.GetVillageByPlayerID(playerID)
    return layouts, err
}