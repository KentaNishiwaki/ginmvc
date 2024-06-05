package controllers

import (
	"ginmvc/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowTop(c *gin.Context, config *models.Config) {
	user, err := models.NewLoginUser(c, config)
	if err != nil {
		c.HTML(http.StatusOK, "Login.html", gin.H{"formdata": models.GetNoUser(config, user)})
		return
	}
}

func MoveTop(c *gin.Context, config *models.Config, user *models.LoginUser) {
	formdata := models.GetAll(user, config)
	c.HTML(http.StatusOK, "Index.html", gin.H{"formdata": formdata})
}

func ExecShutdown(c *gin.Context, config *models.Config) {
	go models.ExecShutdown(config)
	c.Redirect(http.StatusMovedPermanently, "/")
}

func ExecReboot(c *gin.Context, config *models.Config) {
	go models.ExecReboot(config)
	c.Redirect(http.StatusMovedPermanently, "/")
}
func GetKvMicrowave(c *gin.Context, config *models.Config) {
	json := models.GetMicroWaveJSON(config, c.Query("selDate"))
	c.SecureJSON(http.StatusOK, json)
}
func GetHighFrequency(c *gin.Context, config *models.Config) {
	json := models.GetHighFrequencyJSON(config, c.Query("selDate"))
	c.SecureJSON(http.StatusOK, json)
}
