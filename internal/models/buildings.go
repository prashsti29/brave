package models

import "time"

type Building struct {
    ID             string     `gorm:"primaryKey" json:"id"`
    PlayerID       string     `gorm:"not null" json:"player_id"`
    Type           string     `gorm:"not null" json:"type"`
    Name           string     `gorm:"not null" json:"name"`
    Level          int        `gorm:"default:1" json:"level"`
    UpgradePrice   int        `gorm:"default:0" json:"upgrade_price"`
    UpgradeTime    int        `gorm:"default:0" json:"upgrade_time"`
    Currency       string     `json:"currency"`
    IsUpgrading    bool       `gorm:"default:false" json:"is_upgrading"`
    UpgradeEndTime *time.Time `json:"upgrade_end_time"`
    DunbrochLevel  int        `gorm:"default:1" json:"dunbroch_level"`
    MaxAllowed     int        `gorm:"default:1" json:"max_allowed"`
    MaxHealth      int        `gorm:"not null" json:"max_health"`
    CurrentHealth  int        `gorm:"not null" json:"current_health"`
    CreatedAt      time.Time  `json:"created_at"`
}