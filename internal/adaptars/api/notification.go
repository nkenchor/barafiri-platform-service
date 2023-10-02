package api

import (
	"barafiri-platform-service/internal/core/domain/entity"
	"github.com/gin-gonic/gin"
)

func (hdl *HTTPHandler) GetNotificationByRef(c *gin.Context) {
	notification, err := hdl.notificationService.GetNotificationByRef(c.Param("reference"))

	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}

	c.JSON(200, notification)
}
func (hdl *HTTPHandler) GetNotificationByCode(c *gin.Context) {
	notification, err := hdl.notificationService.GetNotificationByCode(c.Param("code"))

	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}

	c.JSON(200, notification)
}

func (hdl *HTTPHandler) GetAllNotifications(c *gin.Context) {
	notifications, err := hdl.notificationService.GetAllNotifications(c.Param("page"))

	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}

	c.JSON(200, notifications)
}
func (hdl *HTTPHandler) CreateNotification(c *gin.Context) {
	body := entity.Notification{}
	_ = c.BindJSON(&body)

	notification, err := hdl.notificationService.CreateNotification(body)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	c.JSON(201, gin.H{"reference": notification})
}
func (hdl *HTTPHandler) UpdateNotification(c *gin.Context) {
	body := entity.Notification{}
	_ = c.BindJSON(&body)
	notification, err := hdl.notificationService.UpdateNotification(c.Param("reference"), body)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	c.JSON(200, gin.H{"reference": notification})
}
func (hdl *HTTPHandler) EnableNotification(c *gin.Context) {
	body := struct {
		Reference string `json:"reference"`
		Enabled   bool   `json:"is_enabled"`
	}{}

	_ = c.BindJSON(&body)
	notification, err := hdl.notificationService.EnableNotification(c.Param("reference"), body.Enabled)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	c.JSON(200, gin.H{"reference": notification})
}
