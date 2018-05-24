package controllers

import (
	"github.com/gin-gonic/gin"
)

// BaseController is
type BaseController struct{}

// SuccessJSON is ok
func (b BaseController) SuccessJSON(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
	})
}

// SuccessJSONData is ok
func (b BaseController) SuccessJSONData(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"success": true,
		"data":    data,
	})
}
