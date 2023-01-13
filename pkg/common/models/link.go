package models

import (
	"time"

	"github.com/google/uuid"
)

type Link struct {
	ID 				uuid.UUID
	UserID 		uuid.UUID `gorm:"type:uuid;default:nill"`
	Password 	string
	Slug 			string
	Link 			string
	Qr 				string
	IsLock 		bool
	CreatedAt time.Time
	UpdatedAt time.Time

	User User
}