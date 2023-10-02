package ports

import "barafiri-platform-service/internal/core/domain/entity"

type CountryRepository interface {
	CreateCountry(entity.Country) (interface{}, error)
	UpdateCountry(ref string, s entity.Country) (interface{}, error)
	EnableCountry(ref string, enabled bool) (interface{}, error)
	GetCountryByCode(CountryCode string) (interface{}, error)
	GetCountryByRef(ref string) (interface{}, error)
	GetAllCountries(page string) (interface{}, error)
}

type CurrencyRepository interface {
	CreateCurrency(entity.Currency) (interface{}, error)
	UpdateCurrency(ref string, s entity.Currency) (interface{}, error)
	EnableCurrency(ref string, enabled bool) (interface{}, error)
	GetCurrencyByCode(code string) (interface{}, error)
	GetCurrencyByRef(ref string) (interface{}, error)
	GetAllCurrencies(page string) (interface{}, error)
}
type NotificationRepository interface {
	CreateNotification(entity.Notification) (interface{}, error)
	UpdateNotification(ref string, s entity.Notification) (interface{}, error)
	EnableNotification(ref string, enabled bool) (interface{}, error)
	GetNotificationByRef(ref string) (interface{}, error)
	GetNotificationByCode(NotificationCode string) (interface{}, error)
	GetAllNotifications(page string) (interface{}, error)
}
type OtpRepository interface {
	CreateOtp(entity.Otp) (interface{}, error)
	UpdateOtp(ref string, s entity.Otp) (interface{}, error)
	EnableOtp(ref string, enabled bool) (interface{}, error)
	GetOtpByRef(ref string) (interface{}, error)
	GetOtpByCode(OtpCode string) (interface{}, error)
	GetAllOtps(page string) (interface{}, error)
}
type CategoryRepository interface {
	CreateCategory(category entity.Category) (interface{}, error)
	UpdateCategory(reference string, category entity.Category) (interface{}, error)
	EnableCategory(reference string, enabled bool) (interface{}, error)
	GetCategoryByRef(reference string) (interface{}, error)
	GetCategoryByName(name string) (interface{}, error)
	GetAllCategories(page string) (interface{}, error)
}
type IndustryRepository interface {
	CreateIndustry(industry entity.Industry) (interface{}, error)
	UpdateIndustry(reference string, industry entity.Industry) (interface{}, error)
	EnableIndustry(reference string, enabled bool) (interface{}, error)
	GetIndustryByRef(reference string) (interface{}, error)
	GetIndustryByName(name string) (interface{}, error)
	GetAllIndustries(page string) (interface{}, error)
}

type TtlRepository interface {
	UpdateTtl(ttl entity.Ttl) (interface{}, error)
	GetTtl() (interface{}, error)
}
type ConfigurationRepository interface {
	SetConfiguration(key string, configuration interface{}) (interface{}, error)
}
