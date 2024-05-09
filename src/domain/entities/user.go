package entities

import (
	"time"
)

type User struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phoneNumber"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CreateUser struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
	Role        string `json:"role"`
}

type UserCreated struct {
	UserId      int64  `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	AccessToken string `json:accessToken`
}

type ResponseCreate[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}
