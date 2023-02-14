package models

import (
	"log"
	"shortlink/pkg/common/resources/link"
	"shortlink/pkg/common/resources/visit_link"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Link struct {
	ID 				string
	UserID 		string `gorm:"type:uuid;default:nill"`
	Qr 				string `gorm:"type:uuid"`
	Password 	string
	Slug 			string
	Link 			string
	IsLock 		bool
	CreatedAt time.Time
	UpdatedAt time.Time

	User User
}

func (l *Link) BeforeCreate(tx *gorm.DB) error {
	l.ID = uuid.NewString()

	return nil
}

func (l *Link) HashPassword() error {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(l.Password), 10)
	if err != nil {
		log.Fatalln(err)
	}

	l.Password = string(hasedPassword)

	return nil
}

func (l *Link) ComparePassword(plainPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(l.Password), []byte(plainPassword)); err != nil {
		return false
	}

	return true
}

func (l *Link) CreateRequest(req *link.CreateRequest) (link *Link) {
	link = &Link{
		Slug:	req.Slug,
		Link:	req.Link,
	}

	return
}

func (l *Link) CreateLinkResponse() (res *link.CreateResponse) {
	data := &link.CreateResponseData{
		Id: l.ID,
		Slug: l.Slug,
		Link: l.Link,
	}

	res = &link.CreateResponse{
		Code: 201,
		Status: "Created",
		Message: "Link Created",
		Data: data,
	}

	return
}

func (l *Link) CreateResponse() (res link.GetAllLinkData) {
	var isPassworded bool

	if l.Password == "" {
		isPassworded = false
	} else {
		isPassworded =true
	}

	res = link.GetAllLinkData{
		ID: l.ID,
    Slug : l.Slug,
    Link: l.Link,
		Qr: l.Qr,
		Password: isPassworded,
    IsLock: l.IsLock,
    CreatedAt: l.CreatedAt,
	}

	return
}

func (l *Link) CreateDetailResponse() (res *link.GetLinkResponse) {
	var isPassworded bool

	if l.Password == "" {
		isPassworded = false
	} else {
		isPassworded =true
	}

	data := &link.GetLinkData{
		ID: l.ID,
    Slug: l.Slug,
    Link: l.Link,
    IsLock: l.IsLock,
		Password: isPassworded,
    CreatedAt: l.CreatedAt,
    UpdatedAt: l.UpdatedAt,
	}
	
	res = &link.GetLinkResponse{
		Code: 200,
    Status: "OK",
    Data: data,
	}

	return
}

func (l *Link) EditLinkResponse() (res *link.CreateResponse) {
	data := &link.CreateResponseData{
		Id: l.ID,
		Slug: l.Slug,
		Link: l.Link,
	}

	res = &link.CreateResponse{
		Code: 200,
		Status: "OK",
		Message: "Link Edited",
		Data: data,
	}

	return
}

func (l *Link) DeleteResponse() (res *link.Response) {
	res = &link.Response{
		Code: 200,
		Status: "OK",
		Message: "Link Deleted",
	}

	return
}

func (l *Link) AddPasswordResponse() (res *link.Response) {
	res = &link.Response{
		Code: 200,
		Status: "OK",
		Message: "Password Set",
	}

	return
}

func (l *Link) EditPasswordResponse() (res *link.Response) {
	res = &link.Response{
		Code: 200,
		Status: "OK",
		Message: "Password Changed",
	}

	return
}

func (l *Link) DeletePasswordResponse() (res *link.Response) {
	res = &link.Response{
		Code: 200,
		Status: "OK",
		Message: "Password removed",
	}

	return
}

func (l *Link) VisitLinkResponse() (res *visit_link.VisitLinkResponse) {
	res = &visit_link.VisitLinkResponse{
		Code: 200,
    Status: "OK",
    Link: l.Link,
	}
	return
}

func (l *Link) VisitlinkPasswordResponse() (res *visit_link.VisitLinkPasswordResponse) {
	res = &visit_link.VisitLinkPasswordResponse{
		Code: 200,
    Status: "OK",
		Password: true,
	}
	return
}