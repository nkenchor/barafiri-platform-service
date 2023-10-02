package api

import (
	"barafiri-platform-service/internal/core/domain/entity"
	"github.com/gin-gonic/gin"
)

func (hdl *HTTPHandler) GetTtl(c *gin.Context) {
	ttls, err := hdl.ttlService.GetTtl()

	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}

	c.JSON(200, ttls)
}

func (hdl *HTTPHandler) UpdateTtl(c *gin.Context) {
	body := entity.Ttl{}
	_ = c.BindJSON(&body)
	ttl, err := hdl.ttlService.UpdateTtl(body)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	c.JSON(201, ttl)
}
