package main

import (
	"fmt"
	"personalities-api/database"
	"personalities-api/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando servidor...")
	routes.HandleRequest()
}
