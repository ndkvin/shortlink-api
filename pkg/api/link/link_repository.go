package link

import (
	"errors"
	"fmt"

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

func (r *Repository) CreateLink(req *link.CreateRequest, userId, qr string) (res *link.CreateResponse, err error) {

	var link *models.Link

	link = link.CreateRequest(req)
	link.UserID = userId
	link.Qr = qr

	if islAvailable := r.isSlugAvailable(link.Slug); !islAvailable {
		err = fiber.NewError(fiber.StatusBadRequest, "Name has been taken")
		return
	}

	result := r.Db.Create(&link)

	if result.Error != nil {
		err = fiber.ErrInternalServerError

		return
	}

	res = link.CreateLinkResponse()
	return
}

func (r *Repository) GetAllLink(userId string) (res *link.GetAllLinkResponse ,err error) {
	var links []models.Link

	if err = r.Db.Order("created_at desc").Where("user_id = ?", userId).Find(&links).Error; err != nil {
		return
	}

	res  = &link.GetAllLinkResponse{
		Code: 200,
		Status: "OK",
		Data: make([]link.GetAllLinkData, len(links)),
	}

	for i := range(links) {
		res.Data[i] = links[i].CreateResponse()
	}

	return
}

func (r *Repository) getLink(id, userId string) (link *models.Link, err error) {
	if err = r.Db.Where("id = ? AND user_id = ?", id, userId).First(&link).Error; err != nil {
		err = fiber.NewError(fiber.StatusNotFound, "Link Not Found")
		return
	}

	return
}

func (r *Repository) GetLink(id, userId string) (res *link.GetLinkResponse, err error) {
	link, err := r.getLink(id,userId)
	if  err != nil {
		return
	}

	res = link.CreateDetailResponse()
	
	return
}

func (r *Repository) EditLink(req *link.CreateRequest, id, userId, qr string) (res *link.CreateResponse, err error) {
	link, err := r.getLink(id, userId)

	if  err != nil {
		return
	}
	fmt.Println(link.Qr)

	if err = DeleteQR(link.Qr); err !=  nil {
		return
	}

	if islAvailable := r.isSlugAvailable(req.Slug); !islAvailable && link.Slug != req.Slug  {
		err = fiber.NewError(fiber.StatusBadRequest, "Name has been taken")
		return
	}

	link.Link = req.Link
	link.Slug = req.Slug
	link.Qr = qr

	if result := r.Db.Save(link); result.Error != nil {
		err = fiber.ErrInternalServerError
		return
	}

	res = link.EditLinkResponse()
	return 
}

func (r *Repository) DeleteLink(id, UserId string) (res *link.Response, err error) {
	link, err :=r.getLink(id, UserId)

	if err != nil {
		return 
	}
	
	if result := r.Db.Where("user_id =?", UserId).Delete(&link); result.Error != nil {
		err = fiber.ErrInternalServerError
		return
	}
	
	res = link.DeleteResponse()
	return
}

func (r *Repository) AddPassword(id, password, userId string) (res *link.Response, err error) {
	link, err := r.getLink(id, userId)
	if err != nil {
		return
	}

	if link.Password != "" {
		err = fiber.NewError(fiber.StatusForbidden, "No Access")
		return 
	}

	link.Password = password

	link.HashPassword()

	if result := r.Db.Save(link); result.Error != nil {
		err = fiber.ErrInternalServerError
		return
	}

	res = link.AddPasswordResponse()
	return 
}

func (r *Repository) editPassword(link *models.Link, newPassword string) (err error) {
	link.Password  = newPassword

	if newPassword == "" {
		r.Db.Save(link)
		return
	}

	link.HashPassword()

	if res := r.Db.Save(link); res.Error != nil {
		return fiber.ErrInternalServerError
	}

	return
}

func (r *Repository) EditPassword(req *link.EditPasswordRequest, id, userId string) (res *link.Response, err error) {
	link, err := r.getLink(id, userId)
	if err != nil {
		return
	}

	if result := link.ComparePassword(req.OldPassword); !result {
		err = fiber.NewError(fiber.StatusBadRequest, "Old password not match")
		return
	}

	if err = r.editPassword(link, req.NewPassword); err != nil {
		return
	}

	res = link.EditPasswordResponse()
	return 
}

func (r *Repository) DeletePassword(id, password, userId string) (res *link.Response, err error) {
	link, err := r.getLink(id, userId)
	if err != nil {
		return
	}

	if result := link.ComparePassword(password); !result {
		err = fiber.NewError(fiber.StatusBadRequest, "Old password not match")
		return
	}

	if err = r.editPassword(link, ""); err != nil {
		return
	}

	res = link.DeletePasswordResponse()
	return 
}