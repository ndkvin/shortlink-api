package user

import (
	"log"

	"shortlink/pkg/common/models"

	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func (h *Repository) CreateUser(req *CreateRequest) (UserId uint, err error){
	// default return value
	UserId = 0
	err = nil
	
	var user models.User

	user.Email = req.Email
	user.Name = req.Name
	user.Password = req.Password

	result := h.Db.Create(&user)

	if result.Error != nil {
		log.Fatalln(result.Error)
		err = result.Error
		return
	}

	UserId = user.ID
	return
}