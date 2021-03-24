package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ikeyu0806/movie-info-backend/models"
)

type ReviewController struct{}

func (pc ReviewController) Create(c *gin.Context) {
	fmt.Println("exec review create function")

	var r models.ReviewModel
	result, err := r.CreateReview(c)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, result)
	} else {
		c.JSON(500, err)
	}
	c.JSON(200, result)
}

func (pc ReviewController) FindByMovieID(c *gin.Context) {
	fmt.Println("exec review find_by_movie_id function")

	var r models.ReviewModel

	result, err := r.FindByMovieId(c)
	fmt.Println(err)
	c.JSON(200, result)
}
