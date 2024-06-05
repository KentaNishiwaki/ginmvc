package controllers

import (
	"errors"
	"fmt"
	"ginmvc/app/models"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func ShowNoRoute(c *gin.Context, config *models.Config) {
	formdata := models.GetNoUser(config, &models.LoginUser{})
	formdata.Description = "Page not found"
	formdata.Error = true
	fmt.Printf("ShowNoRoute\n")
	c.HTML(200, "404.html", gin.H{"formdata": formdata})
}

func ShowNoMethod(c *gin.Context, config *models.Config) {
	formdata := models.GetNoUser(config, &models.LoginUser{})
	formdata.Description = "Method not allowed"
	formdata.Error = true
	fmt.Printf("ShowNoMethod\n")
	c.HTML(200, "404.html", gin.H{"formdata": formdata})
}

func newJwtMiddleware(config *models.Config) (*jwt.GinJWTMiddleware, error) {

	jwtMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            config.Jwt.Realm,
		Key:              []byte(config.Jwt.Key),
		Timeout:          time.Minute * time.Duration(config.Jwt.Timeout),
		MaxRefresh:       time.Minute * time.Duration(config.Jwt.MaxRefresh),
		SendCookie:       config.Jwt.SendCookie,
		TokenLookup:      config.Jwt.TokenLookup,
		SigningAlgorithm: config.Jwt.SigningAlgorithm,
		TokenHeadName:    config.Jwt.TokenHeadName,
		CookieName:       config.Jwt.CookieName,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			return jwt.MapClaims{
				jwt.IdentityKey: data,
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var l models.LoginUser
			if err := c.ShouldBind(&l); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if !l.IsValid(config) {
				return "", jwt.ErrFailedAuthentication
			}
			return l, nil
		},
		LoginResponse: func(c *gin.Context, code int, tokenstring string, time time.Time) {
		},
	})

	if err != nil {
		return nil, err
	}

	err = jwtMiddleware.MiddlewareInit()
	if err != nil {
		return nil, err
	}

	return jwtMiddleware, nil
}
func handlerMiddleWare(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(context *gin.Context) {
		errInit := authMiddleware.MiddlewareInit()
		if errInit != nil {
			log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
		}
	}
}

func MiddleWare(c *gin.Context, config *models.Config) {
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

func SetConfig(fn func(*gin.Context, *models.Config), config *models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		fn(c, config)
	}
}
func SetConfigJWT(fn func(*gin.Context, *models.Config, *jwt.GinJWTMiddleware), config *models.Config, jwt *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		fn(c, config, jwt)
	}
}
func SetConfigUser(fn func(*gin.Context, *models.Config, *models.LoginUser), config *models.Config, user *models.LoginUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		fn(c, config, user)
	}
}

func IsHaveToken(c *gin.Context, jwtm *jwt.GinJWTMiddleware) (l *models.LoginUser, err error) {
	strtoken, err := c.Cookie("jwt")
	if err != nil {
		return &models.LoginUser{}, errors.New("No Cookie")
	}
	if strtoken == "" {
		return &models.LoginUser{}, errors.New("No Token")
	}

	token, err := jwtm.ParseTokenString(strtoken)
	if err != nil {
		return &models.LoginUser{}, errors.New("Paese Err")
	}
	claims := jwt.ExtractClaimsFromToken(token)
	var indenttity = claims[jwt.IdentityKey]
	var user = models.JsonToLoginUser(indenttity)
	return &user, nil

}
