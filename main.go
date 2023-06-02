package main

import (
	"fmt"
	"personalities-api/models"
	"personalities-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{1, "nome1", "historia 1"},
		{2, "nome2", "historia 2"},
		{3, "nome3", "historia 3"},
	}

	fmt.Println("Iniciando servidor...")
	routes.HandleRequest()
}
