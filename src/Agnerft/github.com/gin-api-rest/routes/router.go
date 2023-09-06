package routes

import (
	"github.com/gin-api-rest/controller"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/alunos", controller.ExibeAlunos)
	r.GET("/alunos/:ID", controller.ExibeUmAluno)
	r.GET("/:nome", controller.Saudacao)
	r.DELETE("/alunos/:ID", controller.DeletarUmAluno)
	r.POST("/alunos", controller.CadastrandoAluno)
	r.Run(":8000")
}
