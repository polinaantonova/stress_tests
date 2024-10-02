package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"postgres/internal/ammoPreparation"
	httpSource "postgres/internal/sources/http"
	"postgres/internal/stressTest"
	"time"
	//"postgres/internal/stressTest"
	//"time"
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

	ammo, err := ammoPreparation.PrepareAmmo(db)
	if err != nil {
		log.Fatal(err)
	}
	//postgresTable := postgres.NewPostgres(db)

	//for delay := 500; delay >= 1; delay = delay * 9 / 10 {
	//	fmt.Printf("Query every %v microseconds\n", delay)
	//	err = stressTest.StressTest(postgresTable, ammo, time.Duration(delay), 10000)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	source := httpSource.NewHttpSource("http://localhost:8080/dictionary")

	for delay := 500; delay >= 1; delay = delay * 9 / 10 {
		fmt.Printf("Query every %v microseconds\n", delay)
		err := stressTest.StressTest(source, ammo, time.Duration(delay), 10000)
		if err != nil {
			log.Fatal(err)
		}
	}

}
