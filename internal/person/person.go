package person

import "time"

type Person struct {
	//Id         int       `db:"id" json:"id"`
	Name       string    `db:"name" json:"name"`
	DOB        time.Time `db:"dob" json:"dob"`
	Occupation string    `db:"occupation" json:"occupation"`
	Salary     int       `db:"salary" json:"salary"`
	Address    string    `db:"address" json:"address"`
	City       string    `db:"city" json:"city"`
}

func NewPerson() Person {
	return Person{}
}
