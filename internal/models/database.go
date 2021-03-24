package models

import "gorm.io/gorm"

type database struct {
	db *gorm.DB
}

func New(d *gorm.DB) *database {
	return &database{db: d}
}
