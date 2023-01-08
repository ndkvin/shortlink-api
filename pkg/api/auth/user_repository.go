package auth

import (
	"log"

	"shortlink/pkg/common/models"

	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(Db *gorm.DB) *Repository{
	return &Repository{
		Db: Db,
	}
}

func (h *Repository) CreateUser(req *CreateRequest) (s *CreateResponseSuccess, e *CreateResponseError, err error) {

	var user models.User

	user.Email = req.Email
	user.Name = req.Name

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err!= nil {
		panic(err)
	}

	user.Password = string(hasedPassword)

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
