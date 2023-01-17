package link

import (
	"time"
)

type CreateRequest struct {
	Slug 			string 	`json:"slug" validate:"required,alphaunicode,min=3"`
	Link			string 	`json:"link" validate:"required,uri"`
}

type CreateResponse struct {
	Code			int 								`json:"code"`
	Status 		string 							`json:"status"`
	Message 	string 							`json:"message"`
	Data 			*CreateResponseData `json:"data"`
}

type CreateResponseData struct {
	Id 				string 	`json:"id"`
	Slug 			string 	`json:"slug"`
	Link			string 	`json:"link"`
}

type GetAllLinkData struct {
	ID 				string 		`json:"id"`
	Slug 			string 		`json:"slug"`
	Link 			string 		`json:"link"`
	Password	bool 			`json:"password"`
	IsLock 		bool			`json:"is_lock"`
	CreatedAt time.Time `json:"created_at"`
}

type GetAllLinkResponse struct {
	Code			int 							`json:"code"`
	Status 		string 						`json:"status"`
	Data 			[]GetAllLinkData 	`json:"data"`
}

type GetLinkResponse struct {
	Code			int 					`json:"code"`
	Status 		string 				`json:"status"`
	Data 			*GetLinkData 	`json:"data"`
}

type GetLinkData struct {
	ID 				string 		`json:"id"`
	Slug 			string 		`json:"slug"`
	Link 			string 		`json:"link"`
	Password	bool 			`json:"password"`
	IsLock 		bool			`json:"is_lock"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Response struct {
	Code				int			`json:"code"`
	Status 			string	`json:"status"`
	Message 		string	`json:"message"`
}

type AddPasswordRequest struct {
	Password	string 	`json:"password" validate:"required,ascii,min=8"`
}

type EditPasswordRequest struct {
	OldPassword	string 	`json:"old_password" validate:"required,ascii,min=8"`
	NewPassword	string 	`json:"new_password" validate:"required,ascii,min=8"`
}