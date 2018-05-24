package controllers

import (
	"mugg/sumo/app/cache"
	"mugg/sumo/app/service"

	"github.com/gin-gonic/gin"
)

type (
	// GoodsGroupController is
	GoodsGroupController struct {
		GoodsGroup service.GoodsGroup
		Cache      appCache.Goods
		Base       BaseController
	}
)

// NewGoodsGroup is
func NewGoodsGroup() *GoodsGroupController {
	return &GoodsGroupController{}
}

// Home is
func (s GoodsGroupController) Home(c *gin.Context) {
	// data, _ := s.Googs.GetAll()
	data, _ := s.GoodsGroup.GetAll()
	s.Base.SuccessJSONData(c, data)
}
