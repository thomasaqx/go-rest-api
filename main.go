package main

import (
	"fmt"

	"github.com/thomasaqx/go-rest-api/database"
	"github.com/thomasaqx/go-rest-api/models"
	"github.com/thomasaqx/go-rest-api/routes"
)

func main() {

	

	
	models.Personalities = []models.Personality{
		{Id: "1", Name: "Nome 1", Description: "Historia 1"},
		{Id: "2", Name: "Nome 2", Description: "Historia 2"},
	}
	
	//conectando com o banco de dados
	database.ConnectDB()
	//subindo a rota e o servidor na porta 8080
	fmt.Println("iniciando o servidor em localhost:8080")
	routes.HandleRequest()
}
