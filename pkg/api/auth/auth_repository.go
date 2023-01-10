package auth

import (
	"errors"
	"log"

	"shortlink/pkg/common/models"
	"shortlink/pkg/common/resources/auth"

	"github.com/gofiber/fiber/v2"
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

func (h *Repository) CareateUser(req *auth.CreateRequest)  (successResponse *auth.CreateResponse,err error) {

	var user *models.User

	user = user.CreateRequest(req)

	if islAvailable := h.isEmailAvailable(user.Email); !islAvailable {
		err = fiber.NewError(fiber.StatusNotFound, "Email has been taken")
		return 
	}

	result := h.Db.Create(&user)

	if result.Error != nil {
		log.Fatalln(result.Error)
		err = fiber.ErrInternalServerError
		
		return 
	}

	successResponse = user.CreateResponse()
	return 
}


func (h *Repository) getUserByEmail(email string) (user *models.User, err error) {
	err = h.Db.Where("email = ?", email).First(&user).Error

	return
}

func (h *Repository) Login(req *auth.LoginRequest) (successResponse *auth.CreateResponse, err error) {

	user, err := h.getUserByEmail(req.Email)

	//email not found
	if err != nil {
		err = fiber.NewError(fiber.StatusNotFound,"User email not found")
		return
	}

	// password not match
	if res := user.ComparePassword(req.Password); !res {
		err = fiber.NewError(fiber.StatusBadRequest, "Password not match")
	}
	
	return
}