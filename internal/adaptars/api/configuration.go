package api

import (
	"barafiri-platform-service/internal/core/domain/entity"
	"barafiri-platform-service/internal/core/helper"
	"encoding/json"
)

func (hdl *HTTPHandler) GetCountries() string {
	countries, e1 := hdl.countryService.GetAllCountries("all")
	if e1 != nil {
		helper.LogEvent("ERROR", helper.ErrorMessage(helper.RedisSetupError, e1.Error()))
	}

	countries_, e2 := json.Marshal(countries)
	if e2 != nil {
		helper.LogEvent("ERROR", helper.ErrorMessage(helper.RedisSetupError, e2.Error()))
	}
	return string(countries_)
}

func (hdl *HTTPHandler) GetCurrencies() string {
	currencies, e1 := hdl.currencyService.GetAllCurrencies("all")
	if e1 != nil {
		helper.LogEvent("ERROR", helper.ErrorMessage(helper.RedisSetupError, e1.Error()))
	}

	currencies_, e2 := json.Marshal(currencies)
	if e2 != nil {
		helper.LogEvent("ERROR", helper.ErrorMessage(helper.RedisSetupError, e2.Error()))
	}
	return string(currencies_)
}

func (hdl *HTTPHandler) GetNotifications() string {
	notifications, e1 := hdl.notificationService.GetAllNotifications("all")
	if e1 != nil {
		helper.LogEvent("ERROR", helper.ErrorMessage(helper.RedisSetupError, e1.Error()))
	}

	notifications_, e2 := json.Marshal(notifications)
	if e2 != nil {
		helper.LogEvent("ERROR", helper.ErrorMessage(helper.RedisSetupError, e2.Error()))
	}
	return string(notifications_)
}
func (hdl *HTTPHandler) GetTtlConfiguration() string {
	ttl, e1 := hdl.ttlService.GetTtl()
	if e1 != nil {
		helper.LogEvent("ERROR", helper.ErrorMessage(helper.RedisSetupError, e1.Error()))
	}

	ttl_, e2 := json.Marshal(ttl)
	if e2 != nil {
		helper.LogEvent("ERROR", helper.ErrorMessage(helper.RedisSetupError, e2.Error()))
	}
	return string(ttl_)
}
func (hdl *HTTPHandler) GetOtp() string {
	otp, e1 := hdl.otpService.GetAllOtps("all")
	if e1 != nil {
		helper.LogEvent("ERROR", helper.ErrorMessage(helper.RedisSetupError, e1.Error()))
	}

	otp_, e2 := json.Marshal(otp)
	if e2 != nil {
		helper.LogEvent("ERROR", helper.ErrorMessage(helper.RedisSetupError, e2.Error()))
	}
	return string(otp_)
}
func (hdl *HTTPHandler) GetCategories() string {
	categories, e1 := hdl.categoryService.GetAllCategories("all")
	if e1 != nil {
		helper.LogEvent("ERROR", helper.ErrorMessage(helper.RedisSetupError, e1.Error()))
	}

	categories_, e2 := json.Marshal(categories)
	if e2 != nil {
		helper.LogEvent("ERROR", helper.ErrorMessage(helper.RedisSetupError, e2.Error()))
	}
	return string(categories_)
}
func (hdl *HTTPHandler) GetIndustries() string {
	industries, e1 := hdl.industryService.GetAllIndustries("all")
	if e1 != nil {
		helper.LogEvent("ERROR", helper.ErrorMessage(helper.RedisSetupError, e1.Error()))
	}

	industries_, e2 := json.Marshal(industries)
	if e2 != nil {
		helper.LogEvent("ERROR", helper.ErrorMessage(helper.RedisSetupError, e2.Error()))
	}
	return string(industries_)
}
func (hdl *HTTPHandler) SetConfiguration() {
	helper.LogEvent("INFO", "Loading Platform Configurations from database...")
	configuration := entity.Configuration{}
	configuration.PlatformCountries = hdl.GetCountries()
	configuration.PlatformCurrencies = hdl.GetCurrencies()
	configuration.PlatformNotificationTypes = hdl.GetNotifications()
	configuration.PlatformOtpTypes = hdl.GetOtp()
	configuration.PlatformCategories = hdl.GetCategories()
	configuration.PlatformIndustries = hdl.GetIndustries()
	configuration.PlatformTtlConfiguration = hdl.GetTtlConfiguration()
	helper.LogEvent("INFO", "Hashing Platform Configurations.")
	_, err := hdl.configurationService.SetConfiguration("platform-configurations", configuration)
	if err != nil {
		helper.LogEvent("ERROR", helper.ErrorMessage(helper.RedisSetupError, err.Error()))
		return
	}
}
