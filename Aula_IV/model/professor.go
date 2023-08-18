package model

import "gorm.io/gorm"

type Professor struct {
	gorm.Model
	Id        int    `gorm:"primarykey"`
	Nome      string `gorm: "type:varchar(60); not null"`
	Titulacao string `gorm: "type:varchar(60), not null`
}
