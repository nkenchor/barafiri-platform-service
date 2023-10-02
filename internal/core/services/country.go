package services

import (
	"barafiri-platform-service/internal/core/domain/entity"
	"barafiri-platform-service/internal/core/helper"
	port "barafiri-platform-service/internal/ports"
	"github.com/google/uuid"
)

type countryService struct {
	countryRepository port.CountryRepository
}

func NewCountry(countryRepository port.CountryRepository) *countryService {
	return &countryService{
		countryRepository: countryRepository,
	}
}

func (service *countryService) CreateCountry(country entity.Country) (interface{}, error) {
	country.Reference = uuid.New().String()
	helper.LogEvent("INFO", "Creating country with reference: "+country.Reference)
	if err := helper.Validate(country); err != nil {
		return nil, err
	}
	return service.countryRepository.CreateCountry(country)
}
func (service *countryService) UpdateCountry(reference string, country entity.Country) (interface{}, error) {
	helper.LogEvent("INFO", "Updating country with reference: "+reference)
	_, err := service.GetCountryByRef(reference)
	country.Reference = reference
	if err != nil {
		return nil, err
	}
	if err := helper.Validate(country); err != nil {
		return nil, err
	}
	return service.countryRepository.UpdateCountry(reference, country)
}
func (service *countryService) EnableCountry(reference string, enabled bool) (interface{}, error) {
	helper.LogEvent("INFO", "Enabling country with reference: "+reference)
	_, err := service.GetCountryByRef(reference)
	if err != nil {
		return nil, err
	}
	return service.countryRepository.EnableCountry(reference, enabled)
}

func (service *countryService) GetCountryByRef(reference string) (interface{}, error) {
	helper.LogEvent("INFO", "Getting country with reference: "+reference)
	country, err := service.countryRepository.GetCountryByRef(reference)
	if err != nil {
		return nil, err
	}
	return country, nil
}
func (service *countryService) GetCountryByCode(code string) (interface{}, error) {
	helper.LogEvent("INFO", "Getting country with code: "+code)
	country, err := service.countryRepository.GetCountryByCode(code)
	if err != nil {
		return nil, err
	}
	return country, nil
}

func (service *countryService) GetAllCountries(page string) (interface{}, error) {
	helper.LogEvent("INFO", "Getting all countries...")
	countries, err := service.countryRepository.GetAllCountries(page)
	if err != nil {
		return nil, err
	}
	return countries, nil
}
