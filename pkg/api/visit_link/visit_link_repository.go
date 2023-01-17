package visit_link

import (
	"shortlink/pkg/common/models"

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

func (r *Repository) getLink(slug string) (link *models.Link, err error) {
	err = r.Db.Where("slug = ? AND is_lock = ?", slug, false).First(&link).Error

	return
}

func (r *Repository) VisitLink(slug, ip string) (res interface{}, err error) {
	link, err := r.getLink(slug)

	if err != nil {
		err = fiber.NewError(fiber.StatusNotFound, "Link Not Found")
		return
	}

	if link.Password != "" {
		res = link.VisitlinkPasswordResponse()
		return
	}

	res = link.VisitLinkResponse()
	return
}