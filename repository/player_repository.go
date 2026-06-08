package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/prashsti29/brave/models"
)

type PlayerRepository struct {
	database *sqlx.DB
}

func NewPlayerRepository(connection *sqlx.DB) *PlayerRepository {
	var playerRepo PlayerRepository
	playerRepo.database = connection
	var result *PlayerRepository
	result = &playerRepo
	return result
}

func (playerRepo *PlayerRepository) CreatePlayer(player *models.Player) error {
	var err error
	query := "INSERT INTO players (id, email, password_hash, dunbroch_level, gems, wisps, embis) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err = playerRepo.database.Exec(query, player.ID, player.Email, player.PasswordHash, player.DunbrochLevel, player.Gems, player.Wisps, player.Embis)
	return err
}

func (playerRepo *PlayerRepository) GetPlayerByID(id string) (*models.Player, error) {
	var err error
	var player models.Player
	query := "SELECT * FROM players WHERE id = $1"
	err = playerRepo.database.Get(&player, query, id)
	var result *models.Player
	result = &player
	return result, err
}
