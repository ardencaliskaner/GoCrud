package model

type User struct {
	ID       int    `json:"Id"`
	Name     string `json:"Name" binding:"required"`
	Email    string `json:"Email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}
