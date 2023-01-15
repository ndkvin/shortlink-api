package models

import (
	"shortlink/pkg/common/resources/link"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Link struct {
	ID 				string
	UserID 		string `gorm:"type:uuid;default:nill"`
	Password 	string
	Slug 			string
	Link 			string
	Qr 				string
	IsLock 		bool
	CreatedAt time.Time
	UpdatedAt time.Time

	User User
}

func (l *Link) BeforeCreate(tx *gorm.DB) error {
	l.ID = uuid.NewString()

	return nil
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
	res = link.GetAllLinkData{
		ID: l.ID,
    Slug : l.Slug,
    Link: l.Link,
    IsLock: l.IsLock,
    CreatedAt: l.CreatedAt,
	}

	return
}