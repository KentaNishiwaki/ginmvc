package models

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type LoginUser struct {
	ID         int    `json:"ID"`
	Username   string `json:"Username"`
	Password   string `form:"inputPassword" json:"password" binding:"required"`
	Department string `json:"Department"`
	Email      string `form:"inputEmail" json:"email" binding:"required"`
}

func NewLoginUser(c *gin.Context, config *Config) (*LoginUser, error) {
	token, err := c.Cookie("jwt")
	userid := token
	if err != nil {
		nouser := &LoginUser{ID: -1}
		return nouser, err
	}
	if userid == "" {
		emptyuser := &LoginUser{ID: -1}
		emptyerr := c.Error(errors.New("emptyuser"))
		return emptyuser, emptyerr
	}
	var id int
	id, _ = strconv.Atoi(userid)
	user := &LoginUser{ID: id}
	user.LoginUserInitiaLize(config)
	return user, nil
}
func SetLoginUser(c *gin.Context, email string, password string, config *Config) (*LoginUser, error) {

	user, err := checkUser(email, password, config)

	return user, err
}

func checkUser(email string, password string, config *Config) (*LoginUser, error) {

	db, err := gorm.Open(mysql.New(config.SetDbConfig()), &gorm.Config{})

	db.AutoMigrate(&LoginUser{})
	user := LoginUser{}
	db.Where("Email = ?", email).First(&user)
	if user.ID == 0 {
		user = LoginUser{ID: -1, Email: email}
		err = errors.New("no user")
		return &user, err
	}

	if user.Password != string(getBinaryBySHA256(password)) {
		user = LoginUser{ID: -1, Email: email, Password: user.Password, Department: string(getBinaryBySHA256(password))}
		err = errors.New("bad password")
		return &user, err
	}

	return &user, err

}

func (user *LoginUser) IsValid(config *Config) bool {
	db, err := gorm.Open(mysql.New(config.SetDbConfig()), &gorm.Config{})
	if err != nil {
		return false
	}
	db.AutoMigrate(&LoginUser{})
	dbuser := LoginUser{}
	db.Where("Email = ?", user.Email).First(&dbuser)
	if dbuser.ID == 0 {
		return false
	}
	if dbuser.Password != string(getBinaryBySHA256(user.Password)) {
		return false
	}
	user.ID = dbuser.ID
	user.Username = dbuser.Username
	user.Department = dbuser.Department
	user.Email = dbuser.Email
	return true
}

func getBinaryBySHA256(s string) string {
	r := sha256.Sum256([]byte(s))
	return hex.EncodeToString(r[:])
}
func (user *LoginUser) LoginUserInitiaLize(config *Config) {

	db, _ := gorm.Open(mysql.New(config.SetDbConfig()), &gorm.Config{})
	db.First(&user, user.ID)
}

func JsonToLoginUser(i interface{}) LoginUser {
	var l LoginUser
	b, err := json.Marshal(i)
	if err != nil {
		l = LoginUser{}
	}
	err = json.Unmarshal(b, &l)
	if err != nil {
		l = LoginUser{}
	}
	return l
}
