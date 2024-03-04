package controllers

import (
	"ginmvc/app/models"

	"github.com/gin-gonic/gin"
)

func ShowTop(c *gin.Context) {
	datas := models.GetAll()
	datas.DevMode = false
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
