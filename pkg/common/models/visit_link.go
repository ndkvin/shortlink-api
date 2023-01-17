package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VisitLink struct {
	ID 				string `gorm:"primaryKey"`
	LinkID 		string `gorm:"type:uuid;default:nill;primaryKey"`
	IP 				string
	CreatedAt time.Time

	Link Link
}

func (v *VisitLink) BeforeCreate(tx *gorm.DB) error {
	v.ID = uuid.NewString()

	return nil
}

func (v *VisitLink) CreateRequest(linkId, ip string) (visit_link *VisitLink) {
	visit_link = &VisitLink{
		IP: ip,
		LinkID: linkId,
	}

	return
}