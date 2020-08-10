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
	UserId              int  `json:"user_id"`
	User                User `gorm:"foreignkey:UserId"`
}

type ReviewModel struct{}

func (m ReviewModel) CreateReview(c *gin.Context) (Review, error) {
	db, err := db.GetDB()

	var r Review

	if err := c.ShouldBindJSON(&r); err != nil {
			fmt.Println(err)
			return r, err
	}
	if err := db.Create(&r).Error; err != nil {
			fmt.Println(err)
			return r, err
	}

	return r, err
}

func (m ReviewModel) FindByMovieId(c *gin.Context) ([]Review, error) {
	db, err := db.GetDB()

	var r []Review

	movie_id := c.Param("movie_id")

	db.Preload("User").Debug().Where("movie_id = ?", movie_id).Find(&r)

	defer db.Close()
	return r, err
}
