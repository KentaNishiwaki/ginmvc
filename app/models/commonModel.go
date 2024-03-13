package models

import (
	"github.com/gin-gonic/gin"
)

type LoginUser struct {
	UserID     string
	Password   string
	Department string
	Email      string
}

func (user *LoginUser) LoginUserInitiaLize() {
	user.UserID = user.Email
	user.Department = "999"
}
func NewLoginUser(c *gin.Context) (*LoginUser, error) {
	userCookie, err := c.Cookie("user")
	userid := userCookie
	if err != nil {
		nouser := &LoginUser{UserID: "-1"}
		return nouser, err
	}
	user := &LoginUser{UserID: userid}
	user.LoginUserInitiaLize()
	return user, nil
}
func SetLoginUser(c *gin.Context, email string, password string) (*LoginUser, error) {

	user, err := checkUser(email, password)
	c.SetCookie("user", user.UserID, 3600, "/", "localhost", false, false)
	return user, err
}

func checkUser(email string, password string) (*LoginUser, error) {

	user := &LoginUser{UserID: "9999", Email: email, Password: password}
	user.LoginUserInitiaLize()
	return user, nil

}
