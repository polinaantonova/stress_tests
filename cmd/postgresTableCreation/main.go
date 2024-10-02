package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"os"
	"strconv"
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
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	MakeCSV()
	MakeCSVBoss()

	query := `CREATE TABLE IF NOT EXISTS persons (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100), 
    dob DATE,
    occupation VARCHAR(100),
    salary INTEGER,
    address VARCHAR(100),
    city VARCHAR(100));`

	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}

	query = `CREATE TABLE IF NOT EXISTS employee_boss (
    id SERIAL PRIMARY KEY,
    employee_id INTEGER,
    boss_id INTEGER);`

	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func MakeCSV() {
	dateFile, err := os.OpenFile("data/date.txt", os.O_RDONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer dateFile.Close()

	dates := make([]string, 0, 100)
	scanner := bufio.NewScanner(dateFile)
	for scanner.Scan() {
		dates = append(dates, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	addressFile, err := os.OpenFile("data/address.txt", os.O_RDONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer addressFile.Close()

	addresses := make([]string, 0, 100)
	scanner = bufio.NewScanner(addressFile)
	for scanner.Scan() {
		addresses = append(addresses, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	cityFile, err := os.OpenFile("data/city_town.txt", os.O_RDONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer cityFile.Close()

	cities := make([]string, 0, 100)
	scanner = bufio.NewScanner(cityFile)
	for scanner.Scan() {
		cities = append(cities, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	nameFile, err := os.OpenFile("data/name.txt", os.O_RDONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer nameFile.Close()

	names := make([]string, 0, 100)
	scanner = bufio.NewScanner(nameFile)
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	occupationFile, err := os.OpenFile("data/occupation.txt", os.O_RDONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer occupationFile.Close()

	occupations := make([]string, 0, 100)
	scanner = bufio.NewScanner(occupationFile)
	for scanner.Scan() {
		occupations = append(occupations, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	tableFile, err := os.Create("data/table1.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer tableFile.Close()

	writer := csv.NewWriter(tableFile)
	defer writer.Flush()

	for i := 0; i < 5000000; i++ {
		randomDate := dates[rand.Intn(len(dates))]
		randomSalary := rand.Intn(1000000)
		randomAddress := addresses[rand.Intn(len(addresses))]
		randomCity := cities[rand.Intn(len(cities))]
		randomName := names[rand.Intn(len(names))]
		randomOccupation := occupations[rand.Intn(len(occupations))]
		data := []string{
			randomName,
			randomDate,
			randomOccupation,
			strconv.Itoa(randomSalary),
			randomAddress,
			randomCity,
		}
		err = writer.Write(data)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func MakeCSVBoss() {
	tableFile, err := os.Create("data/tableBoss.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer tableFile.Close()

	writer := csv.NewWriter(tableFile)
	defer writer.Flush()
	bossId := 10

	for employeeId := 1; ; {
		if employeeId >= 5000000 || bossId >= 5000000 {
			break
		}

		for employeeId < bossId {
			data := []string{strconv.Itoa(employeeId), strconv.Itoa(bossId)}
			writer.Write(data)
			employeeId += 1
		}
		employeeId = bossId + 1
		bossId = employeeId + 9

	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
