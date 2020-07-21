package user

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

type Controller struct{}

func (pc Controller) Create(c *gin.Context) {
	fmt.Println(c)
	// var s user.Service
	// p, err := s.CreateModel(c)

	// if err != nil {
	// 		c.AbortWithStatus(400)
	// 		fmt.Println(err)
	// } else {
	// 		c.JSON(201, p)
	// }
}
