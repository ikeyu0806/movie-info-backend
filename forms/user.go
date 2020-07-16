package forms

type UserSignup struct {
	Name       string `json:"name" binding:"required"`
	EMail      string `json:"birthday" binding:"required"`
	Password   string `json:"gender" binding:"required"`
}
