package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/prashsti29/brave/service"
)

type ConfigController struct {
	configService *service.ConfigService
}

func NewConfigController(configService *service.ConfigService) *ConfigController {
	var configController ConfigController
	configController.configService = configService
	var result *ConfigController
	result = &configController
	return result
}

func (c *ConfigController) GetGameConfigs(responseWriter http.ResponseWriter, request *http.Request) {
	configs, err := c.configService.GetGameConfigs()
	if err != nil {
		http.Error(responseWriter, "Could not fetch configs", http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(configs)
}