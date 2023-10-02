package api

import (
	"barafiri-platform-service/internal/core/domain/entity"
	"github.com/gin-gonic/gin"
)

func (hdl *HTTPHandler) GetCountryByRef(c *gin.Context) {
	country, err := hdl.countryService.GetCountryByRef(c.Param("reference"))

	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}

	c.JSON(200, country)
}
func (hdl *HTTPHandler) GetCountryByCode(c *gin.Context) {
	country, err := hdl.countryService.GetCountryByCode(c.Param("code"))

	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}

	c.JSON(200, country)
}

func (hdl *HTTPHandler) GetAllCountries(c *gin.Context) {
	countries, err := hdl.countryService.GetAllCountries(c.Param("page"))

	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}

	c.JSON(200, countries)
}

func (hdl *HTTPHandler) CreateCountry(c *gin.Context) {
	body := entity.Country{}
	_ = c.BindJSON(&body)

	country, err := hdl.countryService.CreateCountry(body)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	c.JSON(201, gin.H{"reference": country})
}
func (hdl *HTTPHandler) UpdateCountry(c *gin.Context) {
	body := entity.Country{}
	_ = c.BindJSON(&body)
	country, err := hdl.countryService.UpdateCountry(c.Param("reference"), body)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	c.JSON(200, gin.H{"reference": country})
}
func (hdl *HTTPHandler) EnableCountry(c *gin.Context) {
	body := struct {
		Reference string `json:"reference"`
		Enabled   bool   `json:"is_enabled"`
	}{}

	_ = c.BindJSON(&body)
	country, err := hdl.countryService.EnableCountry(c.Param("reference"), body.Enabled)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	c.JSON(200, gin.H{"reference": country})
}
