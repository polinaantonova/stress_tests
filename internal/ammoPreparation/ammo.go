package ammoPreparation

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
)

func PrepareAmmo(db *sqlx.DB) ([]int, error) {
	query := `SELECT id from persons;`

	rows, err := db.Query(query)
	if errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	postgresIds := make([]int, 0, 5000000)

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			panic(err)
		}
		postgresIds = append(postgresIds, id)
	}
	log.Println("ammo prepared")
	return postgresIds, err
}
