package controllers

import (
	"mugg/sumo/app/cache"
	"mugg/sumo/app/service"

	"github.com/gin-gonic/gin"
)

type (
	// GoodsBrandController is
	GoodsBrandController struct {
		GoodsBrand service.GoodsBrand
		Cache      appCache.Goods
		Base       BaseController
	}
)

// NewGoodsBrand is
func NewGoodsBrand() *GoodsBrandController {
	return &GoodsBrandController{}
}

// Home is
func (s GoodsBrandController) Home(c *gin.Context) {
	// data, _ := s.Googs.GetAll()
	data, _ := s.GoodsBrand.GetAll()
	s.Base.SuccessJSONData(c, data)
}
