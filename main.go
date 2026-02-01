package main

import (
	"fmt"

	"github.com/thomasaqx/go-rest-api/database"
	"github.com/thomasaqx/go-rest-api/models"
	"github.com/thomasaqx/go-rest-api/routes"
)

func main() {

	//mock
	models.Personalities = []models.Personality{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	//conectando com o banco de dados
	database.ConnectDB()

	//subindo a rota e o servidor na porta 8080
	fmt.Println("iniciando o servidor em localhost:8080")
	routes.HandleRequest()
}
