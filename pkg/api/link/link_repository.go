package link

import (
	"errors"

	"shortlink/pkg/common/models"
	"shortlink/pkg/common/resources/link"

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

func (r *Repository) isSlugAvailable(slug string) bool {
	var link models.Link

	err := r.Db.Where("slug = ?", slug).First(&link).Error

	return errors.Is(err, gorm.ErrRecordNotFound)
}

func (r *Repository) CreateLink(req *link.CreateRequest, userId string) (response *link.CreateResponse, err error) {

	var link *models.Link

	link = link.CreateRequest(req)
	link.UserID = userId

	if islAvailable := r.isSlugAvailable(link.Slug); !islAvailable {
		err = fiber.NewError(fiber.StatusBadRequest, "Name has been taken")
		return
	}

	result := r.Db.Create(&link)

	if result.Error != nil {
		err = fiber.ErrInternalServerError

		return
	}

	response = link.CreateLinkResponse()
	return
}