package visit_link

import (
	"shortlink/pkg/common/resources/visit_link"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	Db         *gorm.DB
	Repository *Repository
	Validation  *Validation
}

func NewHandler(Db *gorm.DB, 	Repository *Repository, Validation  *Validation) *Handler{
	return &Handler{
		Db: Db,
		Repository: Repository,
		Validation: Validation,
	}
}


func (h *Handler) VisitLink(c *fiber.Ctx) (err error) {
	body := &visit_link.VisitLinkRequest{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err = h.Validation.VisitLinkValidation(body); err != nil {
		return
	}

	res, err := h.Repository.VisitLink(c.Params("slug"), body.IP)
	if err != nil {
		return
	}

	return c.Status(200).JSON(res)
}