package repository

import (
    "github.com/jmoiron/sqlx"
    "github.com/prashsti29/brave/models"
)

type BuildingRepository struct {
    database *sqlx.DB
}

func NewBuildingRepository(connection *sqlx.DB) *BuildingRepository {
    var buildingRepo BuildingRepository
    buildingRepo.database = connection
    var result *BuildingRepository
    result = &buildingRepo
    return result
}

func (buildingRepo *BuildingRepository) CreateBuilding(building *models.Building) error {
    var err error
    query := "INSERT INTO buildings (id, player_id, type, name, level, upgrade_price, upgrade_time, currency, is_upgrading, dunbroch_level, max_allowed, max_health, current_health) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)"
    _, err = buildingRepo.database.Exec(query, building.ID, building.PlayerID, building.Type, building.Name, building.Level, building.UpgradePrice, building.UpgradeTime, building.Currency, building.IsUpgrading, building.DunbrochLevel, building.MaxAllowed, building.MaxHealth, building.CurrentHealth)
    return err
}

func (buildingRepo *BuildingRepository) GetBuildingsByPlayerID(playerID string) ([]models.Building, error) {
    var err error
    var buildings []models.Building
    query := "SELECT * FROM buildings WHERE player_id = $1"
    err = buildingRepo.database.Select(&buildings, query, playerID)
    return buildings, err
}