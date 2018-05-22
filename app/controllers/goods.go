package controllers

import (
	"mugg/sumo/app/service"

	"github.com/gin-gonic/gin"
)

type (
	// GoodsController is
	GoodsController struct {
		Googs service.Goods
	}
)

// NewGoods is
func NewGoods() *GoodsController {
	return &GoodsController{}
}

// Home is
func (s GoodsController) Home(c *gin.Context) {

	data, _ := s.Googs.GetAll()
	c.JSON(200, gin.H{
		"data": data,
	})
}
