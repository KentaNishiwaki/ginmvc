package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")
	r.LoadHTMLGlob("app/views/*/*.html")
	r.GET("/", ShowTop)
	r.POST("/shutdown", ExecShutdown)
	r.POST("/reboot", ExecReboot)
	r.NoRoute(ShowNoRoute)
	r.NoMethod(ShowNoMethod)
	r.Use(MiddleWare)
	return r
}
