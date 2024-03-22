// models/socialMedia.go
package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type SocialMedia struct {
	ID             uint      `gorm:"primary_key" json:"id"`
	Name           string    `json:"name" validate:"required"`
	SocialMediaURL string    `json:"social_media_url" validate:"required"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (s *SocialMedia) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}
