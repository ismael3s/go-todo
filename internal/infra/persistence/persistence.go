package persistence

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func connectToPostgres() *sql.DB {
	dsn := "postgres://postgres:postgres@localhost:5434/postgres?sslmode=disable"

	Db, err := sql.Open("postgres", dsn)

	log.Println("Cone", Db)
	if err != nil {
		panic(err)
	}

	err = Db.Ping()

	if err != nil {
		panic(err)
	}

	return Db

}

func GetPostgresConnection() *sql.DB {
	if Db == nil {
		return connectToPostgres()
	}

	log.Println("asdas")

	log.Println(Db)
	return Db
}
