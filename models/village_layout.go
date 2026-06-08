package models

import "time"

type VillageLayout struct {
	PlayerID   string    `db:"player_id"`
	BuildingID string    `db:"building_id"`
	X          int       `db:"x"`
	Y          int       `db:"y"`
	UpdatedAt  time.Time `db:"updated_at"`
}
