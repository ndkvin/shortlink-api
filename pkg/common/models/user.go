package models

import (
	"log"
	"time"

	"shortlink/pkg/common/resources/auth"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        string
	Email     string
	Name      string
	Password  string
	LastLogin *time.Time
	VerifyAt  *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.NewString()
	u.HashPassword()

	return nil
}

func (u *User) CreateRequest(req *auth.CreateRequest) (user *User) {
	user = &User{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}

	return
}

func (u *User) HashPassword() error {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		log.Fatalln(err)
	}

	u.Password = string(hasedPassword)

	return nil
}

func (u *User) ComparePassword(plainPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainPassword)); err != nil {
		return false
	}

	return true
}

func (u *User) RegisterResponse() (res *auth.CreateResponse) {
	data := &auth.ResponseUserData{
		ID:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	res = &auth.CreateResponse{
		Code:    201,
		Status:  "created",
		Message: "User created",
		Data:    data,
	}

	return
}

func (u *User) LoginResponse(accessToken string) (res *auth.LoginResponse) {
	res = &auth.LoginResponse{
		Code:        200,
		Status:      "OK",
		Message:     "Login Success",
		AccessToken: accessToken,
	}

	return
}

func (u *User) EditPasswordResponse() (res *auth.ChangePasswordResponse) {
	res = &auth.ChangePasswordResponse{
		Code: 200,
		Status: "OK",
		Message: "Edit Password Success",
	}

	return
}

func (u *User) EditProfileResponse() (res *auth.CreateResponse) {
	data := &auth.ResponseUserData{
		ID:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	res = &auth.CreateResponse{
		Code:    200,
		Status:  "OK",
		Message: "User updaetd",
		Data:    data,
	}

	return
}

func (u *User) GetUserResponse() (res *auth.CreateResponse) {
	data := &auth.ResponseUserData{
		ID:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	res = &auth.CreateResponse{
		Code:    200,
		Status:  "OK",
		Message: "Success",
		Data:    data,
	}

	return
}