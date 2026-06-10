package models

import "time"

type VillageLayout struct {
    PlayerID   string    `gorm:"not null" json:"player_id"`
    BuildingID string    `gorm:"primaryKey" json:"building_id"`
    X          int       `gorm:"not null" json:"x"`
    Y          int       `gorm:"not null" json:"y"`
    UpdatedAt  time.Time `json:"updated_at"`
}