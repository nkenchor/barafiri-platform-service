package services

import (
	"barafiri-platform-service/internal/core/helper"
	ports "barafiri-platform-service/internal/ports"
)

type configurationService struct {
	configurationRepository ports.ConfigurationRepository
}

func NewConfiguration(
	configurationRepository ports.ConfigurationRepository,
) *configurationService {
	return &configurationService{
		configurationRepository: configurationRepository,
	}
}
func (service *configurationService) SetConfiguration(key string, configuration interface{}) (interface{}, error) {
	helper.LogEvent("INFO", "Setting platform configurations with key: "+key)
	configurations, err := service.configurationRepository.SetConfiguration(key, configuration)
	if err != nil {
		return nil, err
	}
	return configurations, nil
}
