package model

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name string `json:"name"`
	Logo string `json:"logo"`
	Url  string `json:"url"`
}
