package database

import (
	"github.com/gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBanco() {
	conn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	con, err := gorm.Open(postgres.Open(conn))
	if err != nil {
		panic(err.Error())
	}
	con.AutoMigrate(&models.Aluno{})

	DB = con

}
