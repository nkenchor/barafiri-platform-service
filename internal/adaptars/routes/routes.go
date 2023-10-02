package routes

import (
	"barafiri-platform-service/internal/adaptars/api"
	"barafiri-platform-service/internal/core/helper"
	"barafiri-platform-service/internal/core/middleware"
	"barafiri-platform-service/internal/core/services"
	port "barafiri-platform-service/internal/ports"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	countryRepository port.CountryRepository,
	currencyRepository port.CurrencyRepository,
	notificationRepository port.NotificationRepository,
	otpRepository port.OtpRepository,
	categoryRepository port.CategoryRepository,
	industryRepository port.IndustryRepository,
	ttlRepository port.TtlRepository,
	configurationRepository port.ConfigurationRepository) *gin.Engine {
	router := gin.Default()
	countryService := services.NewCountry(countryRepository)
	currencyService := services.NewCurrency(currencyRepository)
	notificationService := services.NewNotification(notificationRepository)
	otpService := services.NewOtp(otpRepository)
	categoryService := services.NewCategory(categoryRepository)
	industryService := services.NewIndustry(industryRepository)
	ttlService := services.NewTtl(ttlRepository)
	configurationService := services.NewConfiguration(configurationRepository)

	handler := api.NewHTTPHandler(
		countryService,
		currencyService,
		notificationService,
		otpService,
		categoryService,
		industryService,
		ttlService,
		configurationService)

	helper.LogEvent("INFO", "Configuring Routes!")
	//handler.GetAllConfigurations()

	router.Use(middleware.LogRequest)
	handler.SetConfiguration()
	//router.Use(middleware.SetHeaders)

	router.Group("/platform/countries")
	{
		router.POST("/platform/countries", handler.CreateCountry)
		router.PUT("/platform/countries/:reference", handler.UpdateCountry)
		router.PUT("/platform/countries/:reference/enabled/:enabled", handler.EnableCountry)
		router.GET("/platform/countries/country-code/:code", handler.GetCountryByCode)
		router.GET("/platform/countries/country-reference/:reference", handler.GetCountryByRef)
		router.GET("/platform/countries/page/:page", handler.GetAllCountries)
	}

	router.Group("/platform/currencies")
	{
		router.POST("/platform/currencies", handler.CreateCurrency)
		router.PUT("/platform/currencies/:reference", handler.UpdateCurrency)
		router.PUT("/platform/currencies/:reference/enabled/:enabled", handler.EnableCurrency)
		router.GET("/platform/currencies/currency-code/:code", handler.GetCurrencyByCode)
		router.GET("/platform/currencies/currency-reference/:reference", handler.GetCurrencyByRef)
		router.GET("/platform/currencies/page/:page", handler.GetAllCurrencies)
	}

	router.Group("/platform/notification-types")
	{
		router.POST("/platform/notification-types", handler.CreateNotification)
		router.PUT("/platform/notification-types/:reference", handler.UpdateNotification)
		router.PUT("/platform/notification-types/:reference/enabled/:enabled", handler.EnableNotification)
		router.GET("/platform/notification-types/notification-code/:code", handler.GetNotificationByCode)
		router.GET("/platform/notification-types/notification-reference/:reference", handler.GetNotificationByRef)
		router.GET("/platform/notification-types/page/:page", handler.GetAllNotifications)
	}

	router.Group("/platform/otp-types")
	{
		router.POST("/platform/otp-types", handler.CreateOtp)
		router.PUT("/platform/otp-types/:reference", handler.UpdateOtp)
		router.PUT("/platform/otp-types/:reference/enabled/:enabled", handler.EnableOtp)
		router.GET("/platform/otp-types/otp-code/:code", handler.GetOtpByCode)
		router.GET("/platform/otp-types/otp-reference/:reference", handler.GetOtpByRef)
		router.GET("/platform/otp-types/page/:page", handler.GetAllOtps)

	}

	router.Group("/platform/categories")
	{
		router.POST("/platform/categories", handler.CreateCategory)
		router.PUT("/platform/categories/:reference", handler.UpdateCategory)
		router.PUT("/platform/categories/:reference/enabled/:enabled", handler.EnableCategory)
		router.GET("/platform/categories/name/:name", handler.GetCategoryByName)
		router.GET("/platform/categories/category-reference/:reference", handler.GetCategoryByRef)
		router.GET("/platform/categories/page/:page", handler.GetAllCategories)
	}
	router.Group("/platform/industries")
	{
		router.POST("/platform/industries", handler.CreateIndustry)
		router.PUT("/platform/industries/:reference", handler.UpdateIndustry)
		router.PUT("/platform/industries/:reference/enabled/:enabled", handler.EnableIndustry)
		router.GET("/platform/industries/name/:name", handler.GetIndustryByName)
		router.GET("/platform/industries/industries-reference/:reference", handler.GetIndustryByRef)
		router.GET("/platform/industries/page/:page", handler.GetAllIndustries)
	}
	router.Group("/platform/configurations")
	{
		//router.GET("/platform/configurations",handler.SetConfiguration)
		router.POST("/platform/configurations/ttl", handler.UpdateTtl)
		router.GET("/platform/configurations/ttl", handler.GetTtl)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404,
			helper.ErrorMessage(helper.NoResourceError, helper.NoResourceFound))
	})
	return router
}
