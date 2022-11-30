package main

import (
	"fmt"

	"github.com/ivancl4udio/go-rest-api/database"
	"github.com/ivancl4udio/go-rest-api/models"
	"github.com/ivancl4udio/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando servidor REST com GO")
	routes.HandleRequest()
}
