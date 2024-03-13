package controllers

import (
	"ginmvc/app/models"

	"github.com/gin-gonic/gin"
)

func ShowTop(c *gin.Context) {
	user, err := models.NewLoginUser(c)
	if err != nil {
		c.HTML(200, "Login.html", gin.H{"datas": models.GetNoUser()})
		return
	}
	datas := models.GetAll(user)
	c.HTML(200, "Index.html", gin.H{"datas": datas})
}

func ExecShutdown(c *gin.Context) {
	go models.ExecShutdown()
	c.Redirect(301, "/")
}

func ExecReboot(c *gin.Context) {
	go models.ExecReboot()
	c.Redirect(301, "/")
}
