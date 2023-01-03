package user

type CreateRequest struct {
	Name 			string `json:"name"`
	Email			string `json:"email"`
	Password	string `json:"password"`
}

type CreateResponse struct {
	Status 	string `json:"status"`
}