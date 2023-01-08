package auth

import (
	"errors"
	"log"

	"shortlink/pkg/common/models"
	"shortlink/pkg/common/resources/auth"

	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(Db *gorm.DB) *Repository {
	return &Repository{
		Db: Db,
	}
}

func (h *Repository) isEmailAvailable(email string) bool {
	var user models.User

	err := h.Db.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}

	return false
}

func (h *Repository) CareateUser(req *auth.CreateRequest)  (sr *auth.CreateResponseSuccess, er *auth.CreateResponseError, c int) {

	var user models.User

	user = user.CreateRequest(req)


	if islAvailable := h.isEmailAvailable(user.Email); !islAvailable {
		c = 400
		er = user.CreateResponseFail("error", "Email has been taken")
		return 
	}

	result := h.Db.Create(&user)

	if result.Error != nil {
		log.Fatalln(result.Error)
		c = 500
		er = user.CreateResponseFail("Internal server error", "An intermal server error occur")
		return
	}

	c = 201
	sr = user.CreateResponseSuccess()
	return 
}
