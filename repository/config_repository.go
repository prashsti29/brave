package repository

import (
	"gorm.io/gorm"
	"github.com/prashsti29/brave/models"
)

type TroopConfigRepository struct {
	database *gorm.DB
}

func NewTroopConfigRepository(db *gorm.DB) *TroopConfigRepository {
	return &TroopConfigRepository{database: db}
}

func (repo *TroopConfigRepository) GetAllTroopConfigs() ([]models.TroopConfig, error) {
	var configs []models.TroopConfig
	result := repo.database.Find(&configs)
	return configs, result.Error
}

func (repo *TroopConfigRepository) GetTroopConfigByID(id string) (*models.TroopConfig, error) {
	var config models.TroopConfig
	result := repo.database.First(&config, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &config, nil
}

func (repo *TroopConfigRepository) GetTroopConfigByName(name string) (*models.TroopConfig, error) {
	var config models.TroopConfig
	result := repo.database.First(&config, "name = ?", name)
	if result.Error != nil {
		return nil, result.Error
	}
	return &config, nil
}

type ProducerBuildingConfigRepository struct {
	database *gorm.DB
}

func NewProducerBuildingConfigRepository(db *gorm.DB) *ProducerBuildingConfigRepository {
	return &ProducerBuildingConfigRepository{database: db}
}

func (repo *ProducerBuildingConfigRepository) GetAllProducerConfigs() ([]models.ProducerBuildingConfig, error) {
	var configs []models.ProducerBuildingConfig
	result := repo.database.Find(&configs)
	return configs, result.Error
}

func (repo *ProducerBuildingConfigRepository) GetProducerConfigByID(id string) (*models.ProducerBuildingConfig, error) {
	var config models.ProducerBuildingConfig
	result := repo.database.First(&config, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &config, nil
}

func (repo *ProducerBuildingConfigRepository) GetProducerConfigByResourceType(resourceType string) (*models.ProducerBuildingConfig, error) {
	var config models.ProducerBuildingConfig
	result := repo.database.First(&config, "resource_type = ?", resourceType)
	if result.Error != nil {
		return nil, result.Error
	}
	return &config, nil
}

type StorageBuildingConfigRepository struct {
	database *gorm.DB
}

func NewStorageBuildingConfigRepository(db *gorm.DB) *StorageBuildingConfigRepository {
	return &StorageBuildingConfigRepository{database: db}
}

func (repo *StorageBuildingConfigRepository) GetAllStorageConfigs() ([]models.StorageBuildingConfig, error) {
	var configs []models.StorageBuildingConfig
	result := repo.database.Find(&configs)
	return configs, result.Error
}

func (repo *StorageBuildingConfigRepository) GetStorageConfigByID(id string) (*models.StorageBuildingConfig, error) {
	var config models.StorageBuildingConfig
	result := repo.database.First(&config, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &config, nil
}

func (repo *StorageBuildingConfigRepository) GetStorageConfigByResourceType(resourceType string) (*models.StorageBuildingConfig, error) {
	var config models.StorageBuildingConfig
	result := repo.database.First(&config, "resource_type = ?", resourceType)
	if result.Error != nil {
		return nil, result.Error
	}
	return &config, nil
}

type DefenseBuildingConfigRepository struct {
	database *gorm.DB
}

func NewDefenseBuildingConfigRepository(db *gorm.DB) *DefenseBuildingConfigRepository {
	return &DefenseBuildingConfigRepository{database: db}
}

func (repo *DefenseBuildingConfigRepository) GetAllDefenseConfigs() ([]models.DefenseBuildingConfig, error) {
	var configs []models.DefenseBuildingConfig
	result := repo.database.Find(&configs)
	return configs, result.Error
}

func (repo *DefenseBuildingConfigRepository) GetDefenseConfigByID(id string) (*models.DefenseBuildingConfig, error) {
	var config models.DefenseBuildingConfig
	result := repo.database.First(&config, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &config, nil
}

func (repo *DefenseBuildingConfigRepository) GetDefenseConfigByName(name string) (*models.DefenseBuildingConfig, error) {
	var config models.DefenseBuildingConfig
	result := repo.database.First(&config, "name = ?", name)
	if result.Error != nil {
		return nil, result.Error
	}
	return &config, nil
}

type LaboratoryConfigRepository struct {
	database *gorm.DB
}

func NewLaboratoryConfigRepository(db *gorm.DB) *LaboratoryConfigRepository {
	return &LaboratoryConfigRepository{database: db}
}

func (repo *LaboratoryConfigRepository) GetAllLaboratoryConfigs() ([]models.LaboratoryConfig, error) {
	var configs []models.LaboratoryConfig
	result := repo.database.Find(&configs)
	return configs, result.Error
}

func (repo *LaboratoryConfigRepository) GetLaboratoryConfigByID(id string) (*models.LaboratoryConfig, error) {
	var config models.LaboratoryConfig
	result := repo.database.First(&config, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &config, nil
}

type ArmyCampConfigRepository struct {
	database *gorm.DB
}

func NewArmyCampConfigRepository(db *gorm.DB) *ArmyCampConfigRepository {
	return &ArmyCampConfigRepository{database: db}
}

func (repo *ArmyCampConfigRepository) GetAllArmyCampConfigs() ([]models.ArmyCampConfig, error) {
	var configs []models.ArmyCampConfig
	result := repo.database.Find(&configs)
	return configs, result.Error
}

func (repo *ArmyCampConfigRepository) GetArmyCampConfigByID(id string) (*models.ArmyCampConfig, error) {
	var config models.ArmyCampConfig
	result := repo.database.First(&config, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &config, nil
}
type BarrackConfigRepository struct {
	database *gorm.DB
}

func NewBarrackConfigRepository(db *gorm.DB) *BarrackConfigRepository {
	return &BarrackConfigRepository{database: db}
}

func (repo *BarrackConfigRepository) GetAllBarrackConfigs() ([]models.BarrackConfig, error) {
	var configs []models.BarrackConfig
	result := repo.database.Find(&configs)
	return configs, result.Error
}

func (repo *BarrackConfigRepository) GetBarrackConfigByID(id string) (*models.BarrackConfig, error) {
	var config models.BarrackConfig
	result := repo.database.First(&config, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &config, nil
}

func (repo *BarrackConfigRepository) GetBarrackConfigByName(name string) (*models.BarrackConfig, error) {
	var config models.BarrackConfig
	result := repo.database.First(&config, "name = ?", name)
	if result.Error != nil {
		return nil, result.Error
	}
	return &config, nil
}
