package services

import (
	"barafiri-platform-service/internal/core/domain/entity"
	"barafiri-platform-service/internal/core/helper"
	port "barafiri-platform-service/internal/ports"
	"github.com/google/uuid"
)

type industryIndustries struct {
	industryRepository port.IndustryRepository
}

func NewIndustry(
	industryRepository port.IndustryRepository,
) *industryIndustries {
	return &industryIndustries{

		industryRepository: industryRepository,
	}
}
func (industry *industryIndustries) CreateIndustry(ind entity.Industry) (interface{}, error) {
	ind.Reference = uuid.New().String()
	helper.LogEvent("INFO", "Creating industry configuration with reference: "+ind.Reference)
	if err := helper.Validate(ind); err != nil {
		return nil, err
	}
	return industry.industryRepository.CreateIndustry(ind)
}

func (industry *industryIndustries) UpdateIndustry(reference string, ind entity.Industry) (interface{}, error) {
	helper.LogEvent("INFO", "Updating industry configuration with reference: "+reference)
	_, err := industry.GetIndustryByRef(reference)
	ind.Reference = reference
	if err != nil {
		return nil, err
	}
	if err := helper.Validate(ind); err != nil {
		return nil, err
	}
	return industry.industryRepository.UpdateIndustry(reference, ind)
}
func (industry *industryIndustries) EnableIndustry(reference string, enabled bool) (interface{}, error) {
	helper.LogEvent("INFO", "Enabling industry configuration with reference: "+reference)
	_, err := industry.GetIndustryByRef(reference)
	if err != nil {
		return nil, err
	}
	return industry.industryRepository.EnableIndustry(reference, enabled)
}

func (industry *industryIndustries) GetIndustryByRef(reference string) (interface{}, error) {
	helper.LogEvent("INFO", "Getting industry configuration with reference: "+reference)
	ind, err := industry.industryRepository.GetIndustryByRef(reference)
	if err != nil {
		return nil, err
	}
	return ind, nil
}
func (industry *industryIndustries) GetIndustryByName(name string) (interface{}, error) {
	helper.LogEvent("INFO", "Getting industry configuration with name: "+name)
	ind, err := industry.industryRepository.GetIndustryByName(name)
	if err != nil {
		return nil, err
	}
	return ind, nil
}

func (industry *industryIndustries) GetAllIndustries(page string) (interface{}, error) {
	helper.LogEvent("INFO", "Getting all industries...")
	ind, err := industry.industryRepository.GetAllIndustries(page)
	if err != nil {
		return nil, err
	}
	return ind, nil
}
