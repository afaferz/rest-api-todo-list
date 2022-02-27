package main

import (
	"fmt"

	"github.com/afaferz/rest-api/db"
	"github.com/afaferz/rest-api/models"
	"github.com/afaferz/rest-api/routes"
)

func main() {

	models.Tasks = []models.Task{
		{Id: 1, Name: "Comprar maça", Status: true, Type: "Compras", Priority: 1, Description: "Comprar maças para salada"},
		{Id: 2, Name: "Lavar a casa", Status: true, Type: "Tarefas domésticas", Priority: 3, Description: "Lavar a casa para festa"},
	}

	db.ConnectDB()
	fmt.Println("Init golang REST Api")
	routes.HandleRequest()
}
