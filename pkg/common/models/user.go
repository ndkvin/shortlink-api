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

func (u *User) CreateRequest(req *auth.CreateRequest) (user *User) {
	user = &User{
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


func (u *User) ComparePassword(plainPassword string) (bool) {
	if err :=bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainPassword)); err != nil {
		return false
	}

	return true
}

func (u *User) CreateResponse() (res *auth.CreateResponse) {
	data := &auth.ResponseUserData{
		ID: u.ID,
		Email: u.Email,
		Name: u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	res = &auth.CreateResponse{
		Code: 201,
		Status:  "created",
		Message: "User created",
		Data: data,
	}

	return
}