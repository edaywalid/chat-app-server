package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Username                    string    `json:"username" gorm:"unique;not null"`
	Password                    string    `json:"password" gorm:"not null"`
	Email                       string    `json:"email" gorm:"unique;not null"`
	PFP_URL                     string    `json:"pfp_url"`
	EmailConfirmationCode       string    `json:"email_confirmation_code"`
	EmailConfirmationCodeExpiry time.Time `json:"email_confirmation_code_expiry" gorm:"default:null"`
	IsVerified                  bool      `json:"is_verified" gorm:"default:false"`
}
