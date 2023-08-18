package database

import (
	"aula4/model"
	"fmt"
)

func Migration() {
	db, err := GetDatabaseConnection()

	if err != nil {
		fmt.Println(err)
		return
	}
	//defer db.Close()
	db.AutoMigrate(&model.Disciplina{})
	db.AutoMigrate(&model.Professor{})
	db.AutoMigrate(&model.Aluno{})
}
