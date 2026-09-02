package models

import "time"

type TroopConfig struct {
	ID                          string    `gorm:"primaryKey" json:"id"`
	Name                        string    `gorm:"unique;not null" json:"name"`
	TrainingTime                int       `gorm:"not null" json:"training_time"`
	DamagePerShot               int       `gorm:"not null" json:"damage_per_shot"`
	Health                      int       `gorm:"not null" json:"health"`
	UnlocksAtDunbrochLevel      int       `json:"unlocks_at_dunbroch_level"`
	Level                       int       `json:"level"`
	HousingSpace                int       `gorm:"not null" json:"housing_space"`
	CostWisps                   int       `gorm:"not null" json:"cost_wisps"`
	CostEmbis                   int       `gorm:"not null" json:"cost_embis"`
	CreatedAt                   time.Time `json:"created_at"`
}

type ProducerBuildingConfig struct {
	ID               string    `gorm:"primaryKey" json:"id"`
	Name             string    `gorm:"unique;not null" json:"name"`
	ResourceType     string    `gorm:"not null" json:"resource_type"`
	ProductionRate   int       `gorm:"not null" json:"production_rate"`
	CreatedAt        time.Time `json:"created_at"`
}

type StorageBuildingConfig struct {
	ID           string    `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"unique;not null" json:"name"`
	ResourceType string    `gorm:"not null" json:"resource_type"`
	MaxStorage   int       `gorm:"not null" json:"max_storage"`
	CreatedAt    time.Time `json:"created_at"`
}

type DefenseBuildingConfig struct {
	ID            string    `gorm:"primaryKey" json:"id"`
	Name          string    `gorm:"unique;not null" json:"name"`
	DamagePerShot int       `gorm:"not null" json:"damage_per_shot"`
	AttackSpeed   float64   `gorm:"not null" json:"attack_speed"`
	TargetType    string    `gorm:"not null" json:"target_type"`
	CreatedAt     time.Time `json:"created_at"`
}

type LaboratoryConfig struct {
	ID               string    `gorm:"primaryKey" json:"id"`
	Name             string    `gorm:"unique;not null" json:"name"`
	MaxResearchLevel int       `gorm:"not null" json:"max_research_level"`
	CreatedAt        time.Time `json:"created_at"`
}

type ArmyCampConfig struct {
	ID              string    `gorm:"primaryKey" json:"id"`
	Name            string    `gorm:"unique;not null" json:"name"`
	TroopCapacity   int       `gorm:"not null" json:"troop_capacity"`
	CreatedAt       time.Time `json:"created_at"`
}

type BarrackConfig struct {
	ID           string    `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"unique;not null" json:"name"`
	TrainingTime int       `gorm:"not null" json:"training_time"`
	CreatedAt    time.Time `json:"created_at"`
}

type BuildingConfig struct {
	ID           string `gorm:"primaryKey" json:"id"`
	Name         string `gorm:"not null" json:"name"`
	Level        int    `gorm:"not null" json:"level"`
	UpgradePrice int    `gorm:"not null" json:"upgrade_price"`
	UpgradeTime  int    `gorm:"not null" json:"upgrade_time"`
	Currency     string `gorm:"not null" json:"currency"`
	MaxHealth    int    `gorm:"not null" json:"max_health"`
	DunbrochLevel int   `gorm:"not null" json:"dunbroch_level"`
	MaxAllowed   int    `gorm:"not null" json:"max_allowed"`
}
