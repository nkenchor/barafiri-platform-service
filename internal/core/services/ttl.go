package services

import (
	"barafiri-platform-service/internal/core/domain/entity"
	"barafiri-platform-service/internal/core/helper"
	ports "barafiri-platform-service/internal/ports"
)

type ttlService struct {
	ttlRepository ports.TtlRepository
}

func NewTtl(
	ttlRepository ports.TtlRepository,
) *ttlService {
	return &ttlService{

		ttlRepository: ttlRepository,
	}
}

func (service *ttlService) UpdateTtl(ttl entity.Ttl) (interface{}, error) {
	helper.LogEvent("INFO", "Updating TTl...")
	if _, err := ttl.TimeUnit.CheckTimeUnitEnum(); err != nil {
		return nil, err
	}
	if err := helper.Validate(ttl); err != nil {
		return nil, err
	}
	return service.ttlRepository.UpdateTtl(ttl)
}

func (service *ttlService) GetTtl() (interface{}, error) {
	helper.LogEvent("INFO", "Getting TTl...")
	ttl, err := service.ttlRepository.GetTtl()
	if err != nil {
		return nil, err
	}
	return ttl, nil
}
