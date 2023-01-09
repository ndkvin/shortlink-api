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

	var user *models.User

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


func (h *Repository) getUserByEmail(email string) (user *models.User, err error) {

	err = h.Db.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	return 
}

func (h *Repository) Login(req *auth.LoginRequest) (sr *auth.CreateResponseSuccess, er *auth.CreateResponseError, c int) {
	var user *models.User
	var err error
	user, err = h.getUserByEmail(req.Email)

	//email not found
	if err != nil {
		er = user.CreateResponseFail("Not Found", "User email not found")
		c = 404

		return
	}

	// password not match
	if res := user.ComparePassword(req.Password); !res {
		er = user.CreateResponseFail("Bad Request", "Password not match in database")
		c = 400
	}
	
	return
}