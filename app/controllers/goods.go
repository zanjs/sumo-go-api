package controllers

import (
	"mugg/sumo/app/cache"
	"mugg/sumo/app/service"

	"github.com/gin-gonic/gin"
)

type (
	// GoodsController is
	GoodsController struct {
		Googs service.Goods
		Cache appCache.Goods
		Base  BaseController
	}
)

// NewGoods is
func NewGoods() *GoodsController {
	return &GoodsController{}
}

// Home is
func (s GoodsController) Home(c *gin.Context) {
	// data, _ := s.Googs.GetAll()
	data, _ := s.Cache.GetAll()
	s.Base.SuccessJSONData(c, data)
}
