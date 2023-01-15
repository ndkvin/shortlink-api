package link

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
	Id 				string `json:"id"`
	Slug 			string 	`json:"name"`
	Link			string 	`json:"email"`
}