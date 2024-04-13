package model

type User struct {
	FirstName   string       `json:"firstName"`
	LastName    string       `json:"lastName"`
	Email       string       `json:"email" binding:"required,email"`
	Password    string       `json:"password"`
	AccessToken *AccessToken `json:"-"`
}
