package inMemory

import (
	"github.com/jmoiron/sqlx"
	"postgres/internal/person"
)

func CreateInMemoryTable(db *sqlx.DB) (map[int]*person.Person, error) {
	query := `SELECT id, name, dob, occupation, salary, address, city FROM persons;`
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	persons := make(map[int]*person.Person, 5000000)
	for rows.Next() {
		var id int
		myPerson := person.NewPerson()
		err = rows.Scan(&id, &myPerson.Name, &myPerson.DOB, &myPerson.Occupation, &myPerson.Salary, &myPerson.Address, &myPerson.City)

		if err != nil {
			return nil, err
		}
		persons[id] = &myPerson
	}
	return persons, nil
}
