package api

import (
	port "barafiri-platform-service/internal/ports"
)

type HTTPHandler struct {
	countryService       port.CountryService
	currencyService      port.CurrencyService
	notificationService  port.NotificationService
	otpService           port.OtpService
	categoryService      port.CategoryService
	industryService      port.IndustryService
	ttlService           port.TtlService
	configurationService port.ConfigurationService
}

func NewHTTPHandler(
	countryService port.CountryService,
	currencyService port.CurrencyService,
	notificationService port.NotificationService,
	otpService port.OtpService,
	categoryService port.CategoryService,
	industryService port.IndustryService,
	ttlService port.TtlService,
	configurationService port.ConfigurationService) *HTTPHandler {
	return &HTTPHandler{
		countryService:       countryService,
		currencyService:      currencyService,
		notificationService:  notificationService,
		otpService:           otpService,
		categoryService:      categoryService,
		industryService:      industryService,
		ttlService:           ttlService,
		configurationService: configurationService,
	}
}
