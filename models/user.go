package models

type User struct {
	ID         int    `json:"id"`
	Email      string `json:"email" gorm:"unique"`
	Password   string `json:"password"`
	OTP        string `json:"otp"`
	IsVerified bool   `json:"is_verified"`
}
