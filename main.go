package main

import (
	"github.com/roberto/go-api-rest/database"
	"github.com/roberto/go-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequestCapiturar()
}
