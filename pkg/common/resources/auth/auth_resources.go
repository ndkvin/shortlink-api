package auth

import (
	"time"

	"github.com/google/uuid"
)

type CreateRequest struct {
	Name 			string 	`json:"name" validate:"required,alphaunicode,min=3,max=16"`
	Email			string 	`json:"email" validate:"required,email"`
	Password	string 	`json:"password" validate:"required,ascii,min=8"`
}

type CreateResponse struct {
	Code			int 							`json:"code"`
	Status 		string 						`json:"status"`
	Message 	string 						`json:"message"`
	Data 			*ResponseUserData `json:"data"`
}

type ResponseUserData struct {
	ID 				uuid.UUID `json:"id"`
	Email     string 		`json:"email"`
	Name			string 		`json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Email			string 	`json:"email" validate:"required,email"`
	Password	string	`json:"password" validate:"required,ascii,min=3,max=16"`
}

type LoginResponse struct {
	Code				int			`json:"code"`
	Status 			string	`json:"status"`
	Message 		string	`json:"message"`
	AccessToken string	`json:"access_token"`
}