package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Disciplina struct {
	gorm.Model
	Nome      string
	Professor string
	Horas     int
	Sala      int
}

func main() {
	var dbUrl string

	//postgres://usuario:senha@ip:porta/banco
	dbUrl = "postgres://postgres:postgres@localhost:5433/aula_orm"
	db, err := gorm.Open(postgres.Open(dbUrl))

	if err != nil {
		fmt.Println(err.Error)
		return
	}

	err = db.AutoMigrate(&Disciplina{})

	if err != nil {
		fmt.Println(err.Error)
		return
	}

	BDII := Disciplina{}
	//Select * FROM Disciplinas ORDER BY ID LIMIT 1
	db.First(&BDII)

	if BDII.ID > 0 {
		if BDII.Sala == 301 {
			fmt.Println("existe o registro e vai apagar!!!!")
			db.Delete(&BDII)
		} else {
			fmt.Println("existe registro")
			BDII.Sala = 301

			//Altera apenas uma coluna
			//update Disciplina SET SALA='301' where id=1
			db.Model(&BDII).Update("Sala", 301)

			//Altera apenas uma coluna
			//update Disciplina SET SALA='301' where id=1
			db.Save(&BDII)

			//Update com filtro (varios registros)
			//update disciplinas SET sala = 301 where sala = 302
			db.Model(&Disciplina{}).Where("sala=?", 302).Update("Sala", 301)

		}

	} else {
		fmt.Println("Realiza inclusão")
		//Inserindo dados no banco
		BDII.Nome = "Banco de Dados II"
		BDII.Horas = 72
		BDII.Professor = "Fabio"
		BDII.Sala = 302

		db.Create(&BDII)
	}

	fmt.Println(BDII)

	fmt.Println("Olá Mundo!")

}
