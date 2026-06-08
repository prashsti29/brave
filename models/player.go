package models

import "time"

type Player struct {
    ID             string    `db:"id"`
    Email          string    `db:"email"`
    PasswordHash   string    `db:"password_hash"`
    CreatedAt      time.Time `db:"created_at"`
    DunbrochLevel  int       `db:"dunbroch_level"`
    Gems           int       `db:"gems"`
    Wisps          int       `db:"wisps"`
    Embis          int       `db:"embis"`
    TotalAttacks   int       `db:"total_attacks"`
    AttacksWon     int       `db:"attacks_won"`
    TotalDefenses  int       `db:"total_defenses"`
    DefensesWon    int       `db:"defenses_won"`
}