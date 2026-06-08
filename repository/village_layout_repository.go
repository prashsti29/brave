package repository

import (
    "github.com/jmoiron/sqlx"
    "github.com/prashsti29/brave/models"
)

type VillageLayoutRepository struct {
    database *sqlx.DB
}

func NewVillageLayoutRepository(connection *sqlx.DB) *VillageLayoutRepository {
    var villageRepo VillageLayoutRepository
    villageRepo.database = connection
    var result *VillageLayoutRepository
    result = &villageRepo
    return result
}

func (villageRepo *VillageLayoutRepository) CreateVillageLayout(layout *models.VillageLayout) error {
    var err error
    query := "INSERT INTO village_layout (player_id, building_id, x, y) VALUES ($1, $2, $3, $4)"
    _, err = villageRepo.database.Exec(query, layout.PlayerID, layout.BuildingID, layout.X, layout.Y)
    return err
}

func (villageRepo *VillageLayoutRepository) GetVillageByPlayerID(playerID string) ([]models.VillageLayout, error) {
    var err error
    var layouts []models.VillageLayout
    query := "SELECT * FROM village_layout WHERE player_id = $1"
    err = villageRepo.database.Select(&layouts, query, playerID)
    return layouts, err
}