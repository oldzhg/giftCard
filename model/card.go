package model

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	Name       string  `json:"name"`
	CategoryId uint    `json:"category_id"`
	Price      float64 `json:"price"`
}
