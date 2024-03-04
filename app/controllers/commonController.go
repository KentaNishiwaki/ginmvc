package controllers

import (
	"fmt"
	"ginmvc/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowNoRoute(c *gin.Context) {
	datas := models.GetAll()
	datas.DevMode = true
	datas.Error = true
	datas.Description = "Page not found"
	c.HTML(200, "404.html", gin.H{"datas": datas})
}

func ShowNoMethod(c *gin.Context) {
	datas := models.GetAll()
	datas.DevMode = true
	datas.Error = true
	datas.Description = "Method not allowed"
	c.HTML(200, "404.html", gin.H{"datas": datas})
}

func MiddleWare(c *gin.Context) {
	fmt.Printf("Total Errors -> %d\n", len(c.Errors))

	if len(c.Errors) <= 0 {
		c.Next()
		return
	}

	for _, err := range c.Errors {
		fmt.Printf("Error -> %+v\n", err)
	}
	c.JSON(http.StatusInternalServerError, "")
}
