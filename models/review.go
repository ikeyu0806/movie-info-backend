package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/ikeyu0806/movie-info-backend/db"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Review struct {
	gorm.Model
	MovieId             int  `json:"movie_id"`
	PublicId            int  `json:"public_id"`
	Comment             string `json:"comment"`
	Score               int  `json:"score"`
}

type ReviewModel struct{}

func (m ReviewModel) CreateReview(c *gin.Context) (Review, error) {
	db := db.GetDB()

	var r Review

	if err := c.ShouldBindJSON(&r); err != nil {
			fmt.Println(err)
			return r, err
	}
	if err := db.Create(&r).Error; err != nil {
			fmt.Println(err)
			return r, err
	}

	return r, nil
}

