package controllers

import (
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"github.com/ikeyu0806/movie-info-backend/models"
)

type UserController struct{}

func (pc UserController) SignUp(c *gin.Context) {
	fmt.Println("exec signup function")

	var s models.UserModel
	s.CreateUser(c)

	secret := "safgvrebwabrq"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin": true,
		"name":   "ikegaya",
		"iat": time.Now(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(201, gin.H{"token": tokenString})
	}
}

func (pc UserController) Login(c *gin.Context) {
	fmt.Println("exec login function")

	var s models.UserModel

	param_name := c.PostForm("name")
	param_password := c.PostForm("password")

	u, err := s.GetByName(param_name)

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	}

	if (u.Name == param_name && u.Password == param_password) {
		fmt.Println("certify")

		secret := "safgvrebwabrq"

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"admin": true,
			"name": "ikegaya",
			"iat": time.Now(),
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, err := token.SignedString([]byte(secret))

		if err != nil {
			c.AbortWithStatus(500)
			fmt.Println(err)
		} else {
			c.JSON(201, gin.H{"user": u, "token": tokenString})
		}
	}
}
