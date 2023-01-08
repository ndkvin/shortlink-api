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
	ID uuid.UUID `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Email string
	Name string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (error) {
 	u.ID = uuid.New()
	u.hashPassword()

  return nil
}

func (u *User) CreateRequest(req *auth.CreateRequest) (user User) {
	user = User{
		Email: req.Email,
		Name: req.Name,
		Password: req.Password,
	}

	return
}

func (u *User) hashPassword() (error) {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		log.Fatalln(err)
	}

	u.Password = string(hasedPassword)

	return nil
}

func (u *User) CreateResponseSuccess() (res *auth.CreateResponseSuccess) {
	data := &auth.ResponseDataUser{
		ID: u.ID,
		Email: u.Email,
		Name: u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	res = &auth.CreateResponseSuccess{
		Status:  "success",
		Message: "User created",
		Data: data,
	}

	return
}

func (u *User) CreateResponseFail(error string, message string) (res *auth.CreateResponseError){
	res = &auth.CreateResponseError{
		Status:  error,
		Message: message,
	}

	return
}