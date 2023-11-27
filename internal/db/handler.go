package db

import "gorm.io/gorm"

type Handler struct {
	db *gorm.DB
}

func NewDb(db *gorm.DB) *Handler {
	return &Handler{db: db}
}
