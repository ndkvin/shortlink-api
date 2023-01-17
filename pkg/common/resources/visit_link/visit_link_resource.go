package visit_link


type VisitLinkRequest struct {
	IP	string	`json:"ip" validate:"required,ipv4"`
}	

type VisitLinkResponse struct {
	Code			int 		`json:"code"`
	Status 		string 	`json:"status"`
	Link			string 	`json:"link"`
}

type VisitLinkPasswordResponse struct {
	Code			int 		`json:"code"`
	Status 		string 	`json:"status"`
	Password	bool 		`json:"password"`
}