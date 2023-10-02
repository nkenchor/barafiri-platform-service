package api

import (
	"barafiri-platform-service/internal/core/domain/entity"
	"github.com/gin-gonic/gin"
)

func (hdl *HTTPHandler) GetCategoryByRef(c *gin.Context) {
	category, err := hdl.categoryService.GetCategoryByRef(c.Param("reference"))

	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}

	c.JSON(200, category)
}
func (hdl *HTTPHandler) GetCategoryByName(c *gin.Context) {
	category, err := hdl.categoryService.GetCategoryByName(c.Param("name"))

	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}

	c.JSON(200, category)
}

func (hdl *HTTPHandler) GetAllCategories(c *gin.Context) {
	categories, err := hdl.categoryService.GetAllCategories(c.Param("page"))

	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}

	c.JSON(200, categories)
}
func (hdl *HTTPHandler) CreateCategory(c *gin.Context) {
	body := entity.Category{}
	_ = c.BindJSON(&body)

	category, err := hdl.categoryService.CreateCategory(body)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	c.JSON(201, gin.H{"reference": category})
}

func (hdl *HTTPHandler) UpdateCategory(c *gin.Context) {
	body := entity.Category{}
	_ = c.BindJSON(&body)
	category, err := hdl.categoryService.UpdateCategory(c.Param("reference"), body)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return

	}

	c.JSON(200, gin.H{"reference": category})
}
func (hdl *HTTPHandler) EnableCategory(c *gin.Context) {
	body := struct {
		Reference string `json:"reference"`
		Enabled   bool   `json:"is_enabled"`
	}{}

	_ = c.BindJSON(&body)
	category, err := hdl.categoryService.EnableCategory(c.Param("reference"), body.Enabled)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	c.JSON(200, gin.H{"reference": category})
}
