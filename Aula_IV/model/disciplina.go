package model

import "gorm.io/gorm"

type Disciplina struct {
	gorm.Model
	Id         int    `gorm:"primarykey"`
	Nome       string `gorm: "type:varchar(60); not null"`
	Atualizado bool   `gorm: "-"`
	//IdProfessor uint        // chave estrangeira para professor
	//Professor   Professores `gorm: "foreignKey: IdProfessor"` //Constroi a relação de chave esrangeira
	ProfessorID uint
	Professor   Professor
	Alunos      []Aluno `gorm: "many to many: disciplina_aluno"`
}
