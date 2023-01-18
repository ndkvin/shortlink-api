package visit_link

import (
	"shortlink/pkg/common/models"
	"shortlink/pkg/common/resources/visit_link"

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

func (r *Repository) addVisitLink(linkId, ip string) (err error) {
	var visit_link *models.VisitLink

	visit_link = visit_link.CreateRequest(linkId, ip)

	if res := r.Db.Create(&visit_link); res.Error != nil {
		return fiber.ErrInternalServerError
	}

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

	if err = r.addVisitLink(link.ID, ip); err != nil {
		return
	}

	res = link.VisitLinkResponse()
	return
}

func (r *Repository) VisitLinkPassword(slug string, body *visit_link.VisitLinkPasswordRequest) (res *visit_link.VisitLinkResponse, err error) {
	link, err := r.getLink(slug)

	if err != nil {
		err = fiber.NewError(fiber.StatusNotFound, "Link Not Found")
		return
	}

	if result := link.ComparePassword(body.Password); !result {
		err = fiber.NewError(fiber.StatusBadRequest, "Password not match")
		return
	}

	res = link.VisitLinkResponse()
	return
}