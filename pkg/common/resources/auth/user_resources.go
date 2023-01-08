package auth

import (
	"time"

	"github.com/google/uuid"
)

type CreateRequest struct {
	Name 			string 	`json:"name" validate:"required,alphaunicode,min=3,max=16"`
	Email			string 	`json:"email" validate:"required,email"`
	Password	string 	`json:"password" validate:"required,alphaunicode,min=8"`
}

type CreateResponseSuccess struct {
	Status 		string 	`json:"status"`
	Message 	string 	`json:"message"`
	Data 			*ResponseDataUser `json:"data"`
}

type ResponseDataUser struct {
	ID 				uuid.UUID `json:"id"`
	Email     string `json:"email"`
	Name			string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateResponseError struct {
	Status 		string 	`json:"status"`
	Message 	string 	`json:"message"`
}