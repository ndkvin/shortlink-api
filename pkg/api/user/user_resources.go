package user

type CreateRequest struct {
	Name 			string 	`json:"name" validate:"required,alphaunicode,min=3,max=16"`
	Email			string 	`json:"email" validate:"required,email"`
	Password	string 	`json:"password" validate:"required,alphaunicode,min=8"`
}

type CreateResponseSuccess struct {
	Status 		string 	`json:"status"`
	Message 	string 	`json:"message"`
	UserId 		uint 		`json:"user_id"`
}

type CreateResponseError struct {
	Status 		string 	`json:"status"`
	Message 	string 	`json:"message"`
}