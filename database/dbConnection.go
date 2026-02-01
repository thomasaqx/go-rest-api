package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {

	fmt.Println("Connection to the postgresql database...")

	//variavel de conexao com o banco de dados
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	//abrindo a conexao com o banco de dados usando funcao Open do gorm
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")

	}
}
