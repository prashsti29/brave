package models

import "time"

type Player struct {
    ID            string    `gorm:"primaryKey" json:"id"`
    Email         string    `gorm:"unique;not null" json:"email"`
    PasswordHash  string    `gorm:"not null" json:"-"`
    CreatedAt     time.Time `json:"created_at"`
    DunbrochLevel int       `gorm:"default:1" json:"dunbroch_level"`
    Gems          int       `gorm:"default:10" json:"gems"`
    Wisps         int       `gorm:"default:500" json:"wisps"`
    Embis         int       `gorm:"default:500" json:"embis"`
    TotalAttacks  int       `gorm:"default:0" json:"total_attacks"`
    AttacksWon    int       `gorm:"default:0" json:"attacks_won"`
    TotalDefenses int       `gorm:"default:0" json:"total_defenses"`
    DefensesWon   int       `gorm:"default:0" json:"defenses_won"`
}