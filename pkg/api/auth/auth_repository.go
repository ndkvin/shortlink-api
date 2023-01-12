package auth

import (
	"errors"
	"time"

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

func (r *Repository) isEmailAvailable(email string) bool {
	var user models.User

	err := r.Db.Where("email = ?", email).First(&user).Error

	return errors.Is(err, gorm.ErrRecordNotFound)
}

func (r *Repository) CareateUser(req *auth.CreateRequest) (successResponse *auth.CreateResponse, err error) {

	var user *models.User

	user = user.CreateRequest(req)

	if islAvailable := r.isEmailAvailable(user.Email); !islAvailable {
		err = fiber.NewError(fiber.StatusBadRequest, "Email has been taken")
		return
	}

	result := r.Db.Create(&user)

	if result.Error != nil {
		err = fiber.ErrInternalServerError

		return
	}

	successResponse = user.CreateRegisterResponse()
	return
}

func (r *Repository) getUserByEmail(email string) (user *models.User, err error) {
	err = r.Db.Where("email = ?", email).First(&user).Error

	return
}

func (r *Repository) updateLogin(user *models.User) (err error) {
	timeNow := time.Now()
	user.LastLogin = &timeNow
	
	if res := r.Db.Save(user); res.Error != nil {
		return fiber.ErrInternalServerError
	}

	return
}

func (r *Repository) Login(req *auth.LoginRequest) (user *models.User, err error) {

	user, err = r.getUserByEmail(req.Email)

	//email not found
	if err != nil {
		err = fiber.NewError(fiber.StatusNotFound, "User email not found")
		return
	}

	// password not match
	if res := user.ComparePassword(req.Password); !res {
		err = fiber.NewError(fiber.StatusBadRequest, "Password not match")
		return
	}

	if err = r.updateLogin(user); err != nil {
		err = fiber.ErrInternalServerError
		return
	}

	return
}

func (r *Repository) getUserById(userId string) (user *models.User, err error) {
	err = r.Db.Where("id = ?", userId).First(&user).Error

	return
}

func (r *Repository) doChangePassword(user *models.User, newPassword string) (err error) {
	user.Password  = newPassword
	user.HashPassword()

	if res := r.Db.Save(user); res.Error != nil {
		return fiber.ErrInternalServerError
	}

	return
}

func (r *Repository) ChangePassword(req *auth.ChangePasswordRequest, userId string) (successResponse *auth.ChangePasswordResponse, err error) {
	user, err := r.getUserById(userId)

	//email not found
	if err != nil {
		err = fiber.NewError(fiber.StatusNotFound, "User id not found")
		return
	}

	// Old password not match
	if res := user.ComparePassword(req.OldPassword); !res {
		err = fiber.NewError(fiber.StatusBadRequest, "Old password not match")
		return
	}

	// do change password
	if err = r.doChangePassword(user, req.NewPassword); err != nil {
		return
	}

	successResponse = user.CreateChangePasswordResponse()
	return
}
