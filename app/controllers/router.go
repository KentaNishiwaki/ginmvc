package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Static("/public", "./public")
	r.LoadHTMLGlob("app/views/*/*.html")

	r.GET("/", func(c *gin.Context) {
		_, err := c.Cookie("user")
		if err != nil {
			fmt.Printf(err.Error() + "\n")
			ShowLogin(c)
		} else {
			ShowTop(c)
			r.POST("/shutdown", ExecShutdown)
			r.POST("/reboot", ExecReboot)
		}
	})
	r.POST("/login", ExecLogin)
	r.NoRoute(ShowNoRoute)
	r.NoMethod(ShowNoMethod)
	r.Use(MiddleWare)
	return r
}
