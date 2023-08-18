package main

import (
	"aula4/api"
	"aula4/database"
)

func main() {
	bootStrap()
}

func bootStrap() {
	database.CreateDatabase() //Criar o banco
	database.Migration()      //Criar as tabelas

	api.InicializeApi() //Sobe a nossa API
}
