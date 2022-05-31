package main

import (
	"fmt"
	"github.com/martadrozsa/curso-golang-web/apigorest/models"
	"github.com/martadrozsa/curso-golang-web/apigorest/routes"
)


func main()  {

	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}
	fmt.Println("Iniciando o servidor Rest com Go")

	routes.HandleRequest()
}
