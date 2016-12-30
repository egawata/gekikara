package models

import (
	"github.com/jinzhu/gorm"
)

type Shop struct {
	gorm.Model
	Name         string `gorm:"type:varchar(255)"`
	Address      string `gorm:"type:varchar(255)"`
	BusinessHour string `gorm:"type:varchar(255)"`
	PostUserId   uint64
}
