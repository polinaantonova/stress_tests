package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"postgres/internal/inMemory"
	"postgres/internal/ping"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "polina"
	password = "1234"
	dbname   = "mydatabase"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sqlx.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(0)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to postgres and ready to shoot")

	persons, err := inMemory.CreateInMemoryTable(db)
	if err != nil {
		log.Fatal(err)
	}

	tableInMemory := inMemory.NewInMemory(persons)

	ping := ping.NewPing()

	http.Handle("/", ping)
	http.Handle("/dictionary", tableInMemory)

	log.Println("running server")

	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
