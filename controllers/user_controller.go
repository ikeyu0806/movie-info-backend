package user

import (
    "fmt"
		"github.com/gin-gonic/gin"
		// "github.com/ikeyu0806/movie-info-backend/models"
		"net/http"
)

type Controller struct{}

func (pc Controller) Create(c *gin.Context) {
	fmt.Println("exec signup function")
	c.JSON(http.StatusOK, "foo")
	// var s models.UserModel
	// p, err := s.CreateUser(c)

	// if err != nil {
	// 		c.AbortWithStatus(400)
	// 		fmt.Println(err)
	// } else {
	// 		c.JSON(201, p)
	// }
}
