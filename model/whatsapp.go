package model

import "gorm.io/gorm"

type WhatsApp struct {
	gorm.Model
	WhatsAppId string `json:"whatsapp_id" gorm:"column:whatsapp_id"`
}

func (WhatsApp) TableName() string {
	return "whatsapps"
}
