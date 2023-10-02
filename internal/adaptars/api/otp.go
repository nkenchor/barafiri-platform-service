package api

import (
	"barafiri-platform-service/internal/core/domain/entity"
	"github.com/gin-gonic/gin"
)

func (hdl *HTTPHandler) GetOtpByRef(c *gin.Context) {
	otp, err := hdl.otpService.GetOtpByRef(c.Param("reference"))

	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}

	c.JSON(200, otp)
}
func (hdl *HTTPHandler) GetOtpByCode(c *gin.Context) {
	otp, err := hdl.otpService.GetOtpByCode(c.Param("code"))

	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}

	c.JSON(200, otp)
}

func (hdl *HTTPHandler) GetAllOtps(c *gin.Context) {
	otps, err := hdl.otpService.GetAllOtps(c.Param("page"))

	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}

	c.JSON(200, otps)
}
func (hdl *HTTPHandler) CreateOtp(c *gin.Context) {
	body := entity.Otp{}
	_ = c.BindJSON(&body)

	otp, err := hdl.otpService.CreateOtp(body)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	c.JSON(201, gin.H{"reference": otp})
}
func (hdl *HTTPHandler) UpdateOtp(c *gin.Context) {
	body := entity.Otp{}
	_ = c.BindJSON(&body)
	otp, err := hdl.otpService.UpdateOtp(c.Param("reference"), body)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	c.JSON(200, gin.H{"reference": otp})
}
func (hdl *HTTPHandler) EnableOtp(c *gin.Context) {
	body := struct {
		Reference string `json:"reference"`
		Enabled   bool   `json:"is_enabled"`
	}{}

	_ = c.BindJSON(&body)
	otp, err := hdl.otpService.EnableOtp(c.Param("reference"), body.Enabled)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	c.JSON(200, gin.H{"reference": otp})
}
