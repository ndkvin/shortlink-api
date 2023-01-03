package user

import (
	"log"
	"shortlink/pkg/common/models"
	"gorm.io/gorm"
	// "github.com/gofiber/fiber/v2"
)

type Repository struct {
	db *gorm.DB
}

func (h *Repository) createUser(req *CreateRequest) (error){
	var user models.User

	user.Email = req.Email
	user.Name = req.Name
	user.Password = req.Password

	result := h.db.Create(&user)

	if result.Error != nil {
		log.Fatalln(result.Error)
		return result.Error
	}

	return nil
}