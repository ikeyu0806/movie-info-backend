package controllers

import (
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"github.com/ikeyu0806/movie-info-backend/models"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

func (pc UserController) SignUp(c *gin.Context) {
	fmt.Println("exec signup function")

	var m models.UserModel
	u, err := m.CreateUser(c)

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
		c.JSON(201, gin.H{"token": tokenString, "user_id": u.ID})
	}
}

func (pc UserController) Login(c *gin.Context) {
	fmt.Println("exec login function")

	var s models.UserModel

	param_name := c.PostForm("name")
	param_password := c.PostForm("password")

	u, err := s.FindByName(param_name)

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	}

	match_pass := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(param_password)) == nil

	if (u.Name == param_name && match_pass) {
		fmt.Println("certify")

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
			c.JSON(201, gin.H{"user": u, "token": tokenString, "user_id": u.ID})
		}
	} else {
		c.JSON(401, gin.H{"user": u})
	}
}

func (pc UserController) FindByName(c *gin.Context) {
	fmt.Println("exec user find_by_movie_name function")

	var u models.UserModel

	result, err := u.FindByName(c.Param("name"))
	fmt.Println(err)
	c.JSON(200, result)
}
