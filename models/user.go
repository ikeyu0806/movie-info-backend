package models

import (
	"errors"
	"github.com/ikeyu0806/movie-info-backend/db"
	"github.com/ikeyu0806/movie-info-backend/forms"
)

// // User is user models property
type User struct {
    ID        uint   `json:"id"`
    PublidId	string `json:"public_id"`
		Name  string `json:"name"`
		Password string `json:"password"`
}

type UserModel struct{}

func (h User) Signup(userPayload forms.UserSignup) (*User, error) {

}
