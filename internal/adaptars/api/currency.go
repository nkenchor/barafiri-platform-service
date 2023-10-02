package api

import (
	"barafiri-platform-service/internal/core/domain/entity"
	"github.com/gin-gonic/gin"
)

func (hdl *HTTPHandler) GetCurrencyByRef(c *gin.Context) {
	currency, err := hdl.currencyService.GetCurrencyByRef(c.Param("reference"))

	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}

	c.JSON(200, currency)
}
func (hdl *HTTPHandler) GetCurrencyByCode(c *gin.Context) {
	currency, err := hdl.currencyService.GetCurrencyByCode(c.Param("code"))

	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}

	c.JSON(200, currency)
}

func (hdl *HTTPHandler) GetAllCurrencies(c *gin.Context) {
	currencies, err := hdl.currencyService.GetAllCurrencies(c.Param("page"))

	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}

	c.JSON(200, currencies)
}
func (hdl *HTTPHandler) CreateCurrency(c *gin.Context) {
	body := entity.Currency{}
	_ = c.BindJSON(&body)

	currency, err := hdl.currencyService.CreateCurrency(body)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	c.JSON(201, gin.H{"reference": currency})
}
func (hdl *HTTPHandler) UpdateCurrency(c *gin.Context) {
	body := entity.Currency{}
	_ = c.BindJSON(&body)
	currency, err := hdl.currencyService.UpdateCurrency(c.Param("reference"), body)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	c.JSON(200, gin.H{"reference": currency})
}

func (hdl *HTTPHandler) EnableCurrency(c *gin.Context) {
	body := struct {
		Reference string `json:"reference"`
		Enabled   bool   `json:"is_enabled"`
	}{}

	_ = c.BindJSON(&body)
	currency, err := hdl.currencyService.EnableCurrency(c.Param("reference"), body.Enabled)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	c.JSON(200, gin.H{"reference": currency})
}
