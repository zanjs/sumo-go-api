package main

import (
	"mugg/sumo/app/conf"
	"mugg/sumo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := conf.Config
	routes.API(r)

	r.Run(":" + config.APP.Port)
}
