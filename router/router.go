package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ward-go/common"
	"ward-go/middleware"
	"ward-go/service"
)

func Router() *gin.Engine {
	engine := gin.New()

	//engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	engine.Use(middleware.Recover())

	engine.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "error/404.gohtml", nil)
	})

	engine.GET("/api/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, service.Info.GetInfo())
	})

	engine.GET("/api/usage", func(c *gin.Context) {
		c.JSON(http.StatusOK, service.Usage.GetUsage())
	})

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/index.gohtml", gin.H{"theme": common.Config.Theme, "info": service.Info.GetInfo()})
	})

	return engine

}
