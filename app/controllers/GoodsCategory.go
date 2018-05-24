package controllers

import (
	"mugg/sumo/app/cache"
	"mugg/sumo/app/service"

	"github.com/gin-gonic/gin"
)

type (
	// GoodsCategoryController is
	GoodsCategoryController struct {
		GoodsCategory service.GoodsCategory
		Cache         appCache.Goods
		Base          BaseController
	}
)

// NewGoodsCategory is
func NewGoodsCategory() *GoodsCategoryController {
	return &GoodsCategoryController{}
}

// Home is
func (s GoodsCategoryController) Home(c *gin.Context) {
	// data, _ := s.Googs.GetAll()
	data, _ := s.GoodsCategory.GetAll()
	s.Base.SuccessJSONData(c, data)
}
