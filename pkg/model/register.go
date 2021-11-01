package model

type Register struct {
	Name     string `json:"Name" binding:"required"`
	Email    string `json:"Email" binding:"required"`
	Password string `json:"Password"  binding:"required"`
}
