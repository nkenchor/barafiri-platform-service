package api

import (
	"barafiri-platform-service/internal/core/domain/entity"
	"github.com/gin-gonic/gin"
)

func (hdl *HTTPHandler) GetIndustryByRef(c *gin.Context) {
	industry, err := hdl.industryService.GetIndustryByRef(c.Param("reference"))

	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}

	c.JSON(200, industry)
}
func (hdl *HTTPHandler) GetIndustryByName(c *gin.Context) {
	industry, err := hdl.industryService.GetIndustryByName(c.Param("name"))

	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}

	c.JSON(200, industry)
}

func (hdl *HTTPHandler) GetAllIndustries(c *gin.Context) {
	industrys, err := hdl.industryService.GetAllIndustries(c.Param("page"))

	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}

	c.JSON(200, industrys)
}
func (hdl *HTTPHandler) CreateIndustry(c *gin.Context) {
	body := entity.Industry{}
	_ = c.BindJSON(&body)

	industry, err := hdl.industryService.CreateIndustry(body)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	c.JSON(201, gin.H{"reference": industry})
}
func (hdl *HTTPHandler) UpdateIndustry(c *gin.Context) {
	body := entity.Industry{}
	_ = c.BindJSON(&body)
	industry, err := hdl.industryService.UpdateIndustry(c.Param("reference"), body)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	c.JSON(200, gin.H{"reference": industry})
}
func (hdl *HTTPHandler) EnableIndustry(c *gin.Context) {
	body := struct {
		Reference string `json:"reference"`
		Enabled   bool   `json:"is_enabled"`
	}{}

	_ = c.BindJSON(&body)
	industry, err := hdl.industryService.EnableIndustry(c.Param("reference"), body.Enabled)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	c.JSON(200, gin.H{"reference": industry})
}
