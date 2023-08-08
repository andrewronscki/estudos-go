package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)


func main() {
	const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "123456"
    dbname   = "golang"
	)

	const connStr = "user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"


	db, err := sql.Open("postgres", connStr)
	if err != nil {
			log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	} 

	fmt.Println("Conexão está aberta!")

	linhas, erro := db.Query("select * from golang.usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}