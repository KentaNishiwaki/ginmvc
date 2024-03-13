package controllers

import (
	"ginmvc/app/models"

	"fmt"

	"github.com/gin-gonic/gin"
)

func ShowLogin(c *gin.Context) {
	fmt.Printf("ShowLogin()\n")
	datas := models.GetNoUser()
	datas.Description = "Please input a email and pass"
	c.HTML(200, "Login.html", gin.H{"datas": datas})
}

func ExecLogin(c *gin.Context) {
	email := c.PostForm("inputEmail")
	password := c.PostForm("inputPassword")
	user, err := models.SetLoginUser(c, email, password)
	if err != nil {
		ShowLogin(c)
		return
	}
	datas := models.GetAll(user)

	c.HTML(200, "Index.html", gin.H{"datas": datas})

}
