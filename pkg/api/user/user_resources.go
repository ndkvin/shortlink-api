package user

type CreateRequest struct {
	Name 			string 	`json:"name"`
	Email			string 	`json:"email"`
	Password	string 	`json:"password"`
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