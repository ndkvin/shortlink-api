package models

import (
	"time"

	"github.com/google/uuid"
)

type VisitLink struct {
	ID 				uuid.UUID
	LinkID 		uuid.UUID `gorm:"type:uuid;default:nill"`
	IP 				string
	CreatedAt time.Time

	Link Link
}