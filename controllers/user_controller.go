package user

import (
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"github.com/ikeyu0806/movie-info-backend/models"
)

type Controller struct{}

func (pc Controller) SignUp(c *gin.Context) {
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

	fmt.Println(token)

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(201, gin.H{"token": tokenString})
	}
}

func (pc Controller) Login(c *gin.Context) {
	fmt.Println("exec login function")

	var s models.UserModel
	u, err := s.GetByName("ikeyu0806")

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(201, gin.H{"user": u})
	}
}
