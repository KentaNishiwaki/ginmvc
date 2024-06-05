package controllers

import (
	"ginmvc/app/models"
	"net/http"

	"fmt"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func ShowLogin(c *gin.Context, config *models.Config, user *models.LoginUser) {
	fmt.Printf("ShowLogin()\n")
	formdata := models.GetNoUser(config, user)
	formdata.Description = "Please input a email and pass"
	formdata.MoreStyles = []string{}
	formdata.MoreScripts = []string{}
	formdata.MoreModule = []string{}
	c.HTML(http.StatusOK, "Login.html", gin.H{"formdata": formdata})
}

func ExecLogin(c *gin.Context, config *models.Config, jwt *jwt.GinJWTMiddleware) {
	email := c.PostForm("inputEmail")
	password := c.PostForm("inputPassword")
	user, err := models.SetLoginUser(c, email, password, config)
	if err != nil {
		ShowLogin(c, config, user)
		return
	}

	jwt.LoginHandler(c)

	MoveTop(c, config, user)
}
