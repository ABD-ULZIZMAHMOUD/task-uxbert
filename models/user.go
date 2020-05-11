package models

type User struct {
	Image    string `json:"image"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
