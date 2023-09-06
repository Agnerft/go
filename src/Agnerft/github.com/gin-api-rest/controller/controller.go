package controller

import (
	"net/http"

	"github.com/gin-api-rest/database"
	"github.com/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

var (
	aluno  models.Aluno
	alunos []models.Aluno
)

func ExibeUmAluno(c *gin.Context) {

	id := c.Params.ByName("ID")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ID = " + id + " Não encontrado, Aluno não existe na base"})
		return
	}
	c.JSON(http.StatusOK, &aluno)

}

func DeletarUmAluno(c *gin.Context) {

	id := c.Params.ByName("ID")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Aluno não encontrado"})
		return
	}

	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, aluno)
}

func AtualizaAluno(c *gin.Context) {

	id := c.Params.ByName("ID")

	database.DB.First(&aluno, id)

	aluno.CPF = c.Query("cpf")
	aluno.Nome = c.Query("nome")
	aluno.RG = c.Query("rg")

}

func ExibeAlunos(c *gin.Context) {

	database.DB.Find(&alunos)

	c.JSON(200, &alunos)

}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")

	c.JSON(200, gin.H{
		"API Diz:": "E ai " + nome + ", tudo certo!?",
	})
}

func CadastrandoAluno(c *gin.Context) {

	if err := c.ShouldBindJSON(&aluno); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusCreated, aluno)

}
