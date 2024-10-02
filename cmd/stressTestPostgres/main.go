package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"postgres/internal/ammoPreparation"
	"postgres/internal/sources/postgres"
	"postgres/internal/stressTest"
	"time"
)

const (
	host     = "192.168.1.63"
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

	ammo, err := ammoPreparation.PrepareAmmo(db)
	if err != nil {
		log.Fatal(err)
	}
	postgresSource := postgres.NewPostgres(db)

	for delay := 500; delay >= 1; delay = delay * 9 / 10 {
		fmt.Printf("Query every %v microseconds\n", delay)
		err = stressTest.StressTest(postgresSource, ammo, time.Duration(delay), 10000)
		if err != nil {
			log.Fatal(err)
		}
	}

}
