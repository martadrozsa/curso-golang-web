package main

import (
	"github.com/martadrozsa/curso-golang-web/gogin/models"
	"github.com/martadrozsa/curso-golang-web/gogin/routes"
)

func main() {

	models.Alunos = []models.Aluno{
		{Nome: "Marta", CPF: "123123123", RG: "14151617"},
		{Nome: "Vitor", CPF: "123123129", RG: "14151619"},
	}
	routes.HandlerRequests()
}
