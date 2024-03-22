// models/comment.go
package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Comment struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    uint      `json:"user_id"`
	PhotoID   uint      `json:"photo_id"`
	Message   string    `json:"message" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Comment) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
