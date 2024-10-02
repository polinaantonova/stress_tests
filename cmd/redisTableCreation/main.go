package main

import (
	"context"
	"encoding/csv"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	person2 "postgres/internal/person"
	"strconv"
	"time"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "1234",
		DB:       0,
	})

	ctx := context.Background()

	file, err := os.Open("internal/data/table1.csv")
	if err != nil {
		log.Fatal("cannot open file", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("error while readling file", err)
	}
	for i, record := range records {
		myPerson := person2.NewPerson()
		myPerson.Name = record[0]
		layout := "2006-01-02"
		myPerson.DOB, err = time.Parse(layout, record[1])
		if err != nil {
			log.Fatal("error while date parsing", err)
		}
		myPerson.Occupation = record[2]
		myPerson.Salary, err = strconv.Atoi(record[3])
		if err != nil {
			log.Fatal("cannot convert salary", err)
		}
		myPerson.Address = record[4]
		myPerson.City = record[5]

		id := i + 1
		key := strconv.Itoa(id)
		err = rdb.HMSet(ctx, key, map[string]interface{}{
			"Name":       myPerson.Name,
			"DOB":        myPerson.DOB,
			"Occupation": myPerson.Occupation,
			"Salary":     myPerson.Salary,
			"Address":    myPerson.Address,
			"City":       myPerson.City,
		}).Err()

		if err != nil {
			log.Fatalf("error while adding to Redis: %v", err)
		}
	}

	log.Println("all persons added successfully")
}
