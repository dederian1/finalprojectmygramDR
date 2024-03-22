// models/photo.go
package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Photo struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Title     string    `json:"title" validate:"required"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url" validate:"required"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Photo) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
