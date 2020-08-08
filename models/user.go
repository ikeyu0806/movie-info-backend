package models

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/ikeyu0806/movie-info-backend/db"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserModel struct{}

func (m UserModel) CreateUser(c *gin.Context) (User, error) {
	db, err := db.GetDB()

	var u User

	if err := c.ShouldBindJSON(&u); err != nil {
			return u, err
	}

	if err := db.Create(&u).Error; err != nil {
			return u, err
	}

	return u, err
}

func (m UserModel) GetByName(name string) (User, error) {
	db, err := db.GetDB()
	var u User

	if err := db.Where("name = ?", name).First(&u).Error; err != nil {
			return u, err
	}

	return u, err
}
