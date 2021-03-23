package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type RunningController struct{}

func (pc RunningController) IsRunning(c *gin.Context) {
	fmt.Println("exec review is_running function")

	c.JSON(200, "movie backend server is running")
}
