package ports

import "barafiri-platform-service/internal/core/domain/entity"

type CountryService interface {
	CreateCountry(entity.Country) (interface{}, error)
	UpdateCountry(reference string, s entity.Country) (interface{}, error)
	EnableCountry(reference string, enabled bool) (interface{}, error)
	GetCountryByCode(code string) (interface{}, error)
	GetCountryByRef(ref string) (interface{}, error)
	GetAllCountries(page string) (interface{}, error)
}

type CurrencyService interface {
	CreateCurrency(entity.Currency) (interface{}, error)
	UpdateCurrency(reference string, s entity.Currency) (interface{}, error)
	EnableCurrency(reference string, enabled bool) (interface{}, error)
	GetCurrencyByCode(code string) (interface{}, error)
	GetCurrencyByRef(ref string) (interface{}, error)
	GetAllCurrencies(page string) (interface{}, error)
}
type NotificationService interface {
	CreateNotification(entity.Notification) (interface{}, error)
	UpdateNotification(reference string, s entity.Notification) (interface{}, error)
	EnableNotification(reference string, enabled bool) (interface{}, error)
	GetNotificationByRef(reference string) (interface{}, error)
	GetNotificationByCode(code string) (interface{}, error)
	GetAllNotifications(page string) (interface{}, error)
}
type OtpService interface {
	CreateOtp(entity.Otp) (interface{}, error)
	UpdateOtp(reference string, s entity.Otp) (interface{}, error)
	EnableOtp(reference string, enabled bool) (interface{}, error)
	GetOtpByRef(reference string) (interface{}, error)
	GetOtpByCode(code string) (interface{}, error)
	GetAllOtps(page string) (interface{}, error)
}
type CategoryService interface {
	CreateCategory(category entity.Category) (interface{}, error)
	UpdateCategory(reference string, category entity.Category) (interface{}, error)
	EnableCategory(reference string, enabled bool) (interface{}, error)
	GetCategoryByRef(reference string) (interface{}, error)
	GetCategoryByName(name string) (interface{}, error)
	GetAllCategories(page string) (interface{}, error)
}
type IndustryService interface {
	CreateIndustry(industry entity.Industry) (interface{}, error)
	UpdateIndustry(reference string, industry entity.Industry) (interface{}, error)
	EnableIndustry(reference string, enabled bool) (interface{}, error)
	GetIndustryByRef(reference string) (interface{}, error)
	GetIndustryByName(name string) (interface{}, error)
	GetAllIndustries(page string) (interface{}, error)
}

type TtlService interface {
	UpdateTtl(ttl entity.Ttl) (interface{}, error)
	GetTtl() (interface{}, error)
}

type ConfigurationService interface {
	SetConfiguration(key string, configuration interface{}) (interface{}, error)
}
