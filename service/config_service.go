package service

import (
	"github.com/prashsti29/brave/models"
	"github.com/prashsti29/brave/repository"
)

type ConfigService struct {
	troopRepo               *repository.TroopConfigRepository
	producerRepo            *repository.ProducerBuildingConfigRepository
	storageRepo             *repository.StorageBuildingConfigRepository
	defenseRepo             *repository.DefenseBuildingConfigRepository
	laboratoryRepo          *repository.LaboratoryConfigRepository
	armyCampRepo            *repository.ArmyCampConfigRepository
	barrackRepo             *repository.BarrackConfigRepository
	buildingConfigRepo      *repository.BuildingConfigRepository
}

func NewConfigService(
	troopRepo *repository.TroopConfigRepository,
	producerRepo *repository.ProducerBuildingConfigRepository,
	storageRepo *repository.StorageBuildingConfigRepository,
	defenseRepo *repository.DefenseBuildingConfigRepository,
	laboratoryRepo *repository.LaboratoryConfigRepository,
	armyCampRepo *repository.ArmyCampConfigRepository,
	barrackRepo *repository.BarrackConfigRepository,
	buildingConfigRepo *repository.BuildingConfigRepository,
) *ConfigService {
	return &ConfigService{
		troopRepo:          troopRepo,
		producerRepo:       producerRepo,
		storageRepo:        storageRepo,
		defenseRepo:        defenseRepo,
		laboratoryRepo:     laboratoryRepo,
		armyCampRepo:       armyCampRepo,
		barrackRepo:        barrackRepo,
		buildingConfigRepo: buildingConfigRepo,
	}
}

type GameConfigs struct {
	TroopConfigs             []models.TroopConfig              `json:"troop_configs"`
	ProducerBuildingConfigs  []models.ProducerBuildingConfig   `json:"producer_building_configs"`
	StorageBuildingConfigs   []models.StorageBuildingConfig    `json:"storage_building_configs"`
	DefenseBuildingConfigs   []models.DefenseBuildingConfig    `json:"defense_building_configs"`
	LaboratoryConfigs        []models.LaboratoryConfig         `json:"laboratory_configs"`
	ArmyCampConfigs          []models.ArmyCampConfig           `json:"army_camp_configs"`
	BarrackConfigs           []models.BarrackConfig            `json:"barrack_configs"`
	BuildingConfigs          []models.BuildingConfig           `json:"building_configs"`
}

func (s *ConfigService) GetGameConfigs() (*GameConfigs, error) {
	troops, err := s.troopRepo.GetAllTroopConfigs()
	if err != nil {
		return nil, err
	}

	producers, err := s.producerRepo.GetAllProducerConfigs()
	if err != nil {
		return nil, err
	}

	storages, err := s.storageRepo.GetAllStorageConfigs()
	if err != nil {
		return nil, err
	}

	defenses, err := s.defenseRepo.GetAllDefenseConfigs()
	if err != nil {
		return nil, err
	}

	laboratories, err := s.laboratoryRepo.GetAllLaboratoryConfigs()
	if err != nil {
		return nil, err
	}

	armyCamps, err := s.armyCampRepo.GetAllArmyCampConfigs()
	if err != nil {
		return nil, err
	}

	barracks, err := s.barrackRepo.GetAllBarrackConfigs()
	if err != nil {
		return nil, err
	}

	buildings, err := s.buildingConfigRepo.GetAllBuildingConfigs()
	if err != nil {
		return nil, err
	}

	return &GameConfigs{
		TroopConfigs:            troops,
		ProducerBuildingConfigs: producers,
		StorageBuildingConfigs:  storages,
		DefenseBuildingConfigs:  defenses,
		LaboratoryConfigs:       laboratories,
		ArmyCampConfigs:         armyCamps,
		BarrackConfigs:          barracks,
		BuildingConfigs:         buildings,
	}, nil
}