package model

import "gorm.io/gorm"

type User struct {
	gorm.Model        // include ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `json:"name"`
	Email      string `json:"email"`
}
