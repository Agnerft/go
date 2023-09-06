package main

import (
	"fmt"

	_ "github.com/gin-api-rest/controller"
	"github.com/gin-api-rest/database"
	"github.com/gin-api-rest/routes"
)

func main() {

	database.ConectaComBanco()

	fmt.Println("Servidor iniciado")
	//models.Alunos = []models.Aluno{
	//	{Nome: "Agner", CPF: "12312312313", RG: "9102306363"},
	//	{Nome: "Cleito", CPF: "25821463258", RG: "9102853658"},
	//	{Nome: "Rogério", CPF: "85474120153", RG: "9145876325"},
	//}

	routes.HandleRequests()

}
