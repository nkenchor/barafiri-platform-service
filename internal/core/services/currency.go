package services

import (
	"barafiri-platform-service/internal/core/domain/entity"
	"barafiri-platform-service/internal/core/helper"
	port "barafiri-platform-service/internal/ports"
	"github.com/google/uuid"
)

type currencyService struct {
	currencyRepository port.CurrencyRepository
}

func NewCurrency(
	currencyRepository port.CurrencyRepository,
) *currencyService {
	return &currencyService{

		currencyRepository: currencyRepository,
	}
}

func (service *currencyService) CreateCurrency(currency entity.Currency) (interface{}, error) {
	currency.Reference = uuid.New().String()
	helper.LogEvent("INFO", "Creating currency with reference: "+currency.Reference)
	if err := helper.Validate(currency); err != nil {
		return nil, err
	}
	return service.currencyRepository.CreateCurrency(currency)
}
func (service *currencyService) UpdateCurrency(reference string, currency entity.Currency) (interface{}, error) {
	helper.LogEvent("INFO", "Updating currency with reference: "+reference)
	_, err := service.GetCurrencyByRef(reference)
	currency.Reference = reference
	if err != nil {
		return nil, err
	}
	if err := helper.Validate(currency); err != nil {
		return nil, err
	}
	return service.currencyRepository.UpdateCurrency(reference, currency)
}
func (service *currencyService) EnableCurrency(reference string, enabled bool) (interface{}, error) {
	helper.LogEvent("INFO", "Enabling currency with reference: "+reference)
	_, err := service.GetCurrencyByRef(reference)
	if err != nil {
		return nil, err
	}
	return service.currencyRepository.EnableCurrency(reference, enabled)
}

func (service *currencyService) GetCurrencyByRef(reference string) (interface{}, error) {
	helper.LogEvent("INFO", "Getting currency with reference: "+reference)
	currency, err := service.currencyRepository.GetCurrencyByRef(reference)
	if err != nil {
		return nil, err
	}
	return currency, nil
}
func (service *currencyService) GetCurrencyByCode(code string) (interface{}, error) {
	helper.LogEvent("INFO", "Getting currency with code: "+code)
	currency, err := service.currencyRepository.GetCurrencyByCode(code)
	if err != nil {
		return nil, err
	}
	return currency, nil
}

func (service *currencyService) GetAllCurrencies(page string) (interface{}, error) {
	helper.LogEvent("INFO", "Getting all currencies...")
	currencies, err := service.currencyRepository.GetAllCurrencies(page)
	if err != nil {
		return nil, err
	}
	return currencies, nil
}
