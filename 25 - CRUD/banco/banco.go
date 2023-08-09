package banco

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "estudos_go"
)

const connStr = "user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	} 

	return db, nil
}