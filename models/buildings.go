package models

import "time"

type Building struct {
    ID             string    `db:"id"`
    PlayerID       string    `db:"player_id"`
    Type           string    `db:"type"`
    Name           string    `db:"name"`
    Level          int       `db:"level"`
    UpgradePrice   int       `db:"upgrade_price"`
    UpgradeTime    int       `db:"upgrade_time"`
    Currency       string    `db:"currency"`
    IsUpgrading    bool      `db:"is_upgrading"`
    UpgradeEndTime *time.Time `db:"upgrade_end_time"`
    DunbrochLevel  int       `db:"dunbroch_level"`
    MaxAllowed     int       `db:"max_allowed"`
    MaxHealth      int       `db:"max_health"`
    CurrentHealth  int       `db:"current_health"`
    CreatedAt      time.Time `db:"created_at"`
}