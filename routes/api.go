package routes

import (
	"mugg/sumo/app/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// API is router api
func API(g *gin.Engine) {
	catController := controllers.NewCat()
	goodsController := controllers.NewGoods()
	goodsGroupController := controllers.NewGoodsGroup()
	goodsBrandController := controllers.NewGoodsBrand()
	goodsCategoryController := controllers.NewGoodsCategory()
	goodsSkuController := controllers.NewGoodsSku()

	api := g.Group("/api")

	v1 := api.Group("/v1")
	{
		v1.GET("/", catController.Home)

		v1.GET("/user/:name/:action", func(c *gin.Context) {
			name := c.Param("name")
			action := c.Param("action")
			message := name + " is " + action
			c.String(http.StatusOK, message)
		})
	}

	goods := v1.Group("/goods")
	{
		goods.GET("", goodsController.Home)
	}

	goodsGroup := v1.Group("/goods/group")
	{
		goodsGroup.GET("", goodsGroupController.Home)
	}

	goodsBrand := v1.Group("/goods/brand")
	{
		goodsBrand.GET("", goodsBrandController.Home)
	}

	goodsCategory := v1.Group("/goods/category")
	{
		goodsCategory.GET("", goodsCategoryController.Home)
	}

	goodsSku := v1.Group("/goods/sku")
	{
		goodsSku.GET("", goodsSkuController.Home)
	}

}
