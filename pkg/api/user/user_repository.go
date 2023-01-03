package user

import (
	"log"

	"shortlink/pkg/common/models"

	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func (h *Repository) CreateUser(req *CreateRequest) (s *CreateResponseSuccess, e *CreateResponseError, err error) {

	var user models.User

	user.Email = req.Email
	user.Name = req.Name
	user.Password = req.Password

	result := h.Db.Create(&user)

	if result.Error != nil {
		log.Fatalln(result.Error)
		s = &CreateResponseSuccess{}

		e = &CreateResponseError{
			Status:  "Internal Server error",
			Message: "An intermal server error",
		}

		err = result.Error
		return
	}

	s = &CreateResponseSuccess{
		Status:  "success",
		Message: "User created",
		UserId:  user.ID,
	}

	e = &CreateResponseError{}
	err = nil
	return
}
