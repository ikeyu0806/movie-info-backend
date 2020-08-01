package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "github.com/ikeyu0806/movie-info-backend/models"
)

type ReviewController struct{}

func (pc ReviewController) Create(c *gin.Context) {
	fmt.Println("exec review create function")
}
