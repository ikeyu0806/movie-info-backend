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
	Password  string `json:"password"`
}

type UserModel struct{}

func (s UserModel) CreateUser(c *gin.Context) (User, error) {
	db := db.GetDB()
	db.Create(&User{Name: "test", Password: "pass"})
	var u User

	// if err := c.BindJSON(&u); err != nil {
	// 		return u, err
	// }

	// if err := db.Create(&u).Error; err != nil {
	// 		return u, err
	// }

	return u, nil
}
