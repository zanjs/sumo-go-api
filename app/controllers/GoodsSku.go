package controllers

import (
	"mugg/sumo/app/cache"
	"mugg/sumo/app/service"

	"github.com/gin-gonic/gin"
)

type (
	// GoodsSkuController is
	GoodsSkuController struct {
		GoodsSku service.GoodsSku
		Cache    appCache.Goods
		Base     BaseController
	}
)

// NewGoodsSku is
func NewGoodsSku() *GoodsSkuController {
	return &GoodsSkuController{}
}

// Home is
func (s GoodsSkuController) Home(c *gin.Context) {
	// data, _ := s.Googs.GetAll()
	data, _ := s.GoodsSku.GetAll()
	s.Base.SuccessJSONData(c, data)
}
