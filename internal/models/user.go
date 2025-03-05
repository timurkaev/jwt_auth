package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var Users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}
