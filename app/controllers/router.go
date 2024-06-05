package controllers

import (
	"ginmvc/app/models"
	"log"

	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Static("/public", "./public")
	r.LoadHTMLGlob("app/views/*/*.html")

	var config models.Config
	config, err := config.LoadConfigForYaml()
	if err != nil {
		log.Fatal(err)
		return r
	}

	jwt, err := newJwtMiddleware(&config)
	if err != nil {
		log.Fatal(err)
		return r
	}

	r.GET("/", func(c *gin.Context) {
		var l, err = IsHaveToken(c, jwt)
		if err != nil {
			ShowLogin(c, &config, &models.LoginUser{})
		} else {
			MoveTop(c, &config, l)
		}

	})
	r.POST("/", SetConfigJWT(ExecLogin, &config, jwt))

	api := r.Group("/api", jwt.MiddlewareFunc())
	{
		api.Use(handlerMiddleWare(jwt))
		api.GET("/getkvMicrowave", SetConfig(GetKvMicrowave, &config))
		api.GET("/getHighFrequency", SetConfig(GetHighFrequency, &config))

	}
	r.POST("/reboot", SetConfig(ExecReboot, &config))
	r.GET("/user", SetConfigUser(ShowLogin, &config, &models.LoginUser{}))
	r.POST("/shutdown", SetConfig(ExecShutdown, &config))
	r.NoRoute(SetConfig(ShowNoRoute, &config))
	r.NoMethod(SetConfig(ShowNoMethod, &config))
	r.Use(SetConfig(MiddleWare, &config))
	r.Run(":" + strconv.Itoa(config.App.Port))
	return r
}
