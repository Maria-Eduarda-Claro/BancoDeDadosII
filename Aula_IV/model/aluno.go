package model

type Aluno struct {
	Id   int    `gorm:"primarykey"`
	Nome string `gorm: "type:varchar(100); not null"`
	RA   string `gorm: "type:varchar(60), not null`
}
