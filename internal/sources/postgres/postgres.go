package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"postgres/internal/person"
)

type Postgres struct {
	db *sqlx.DB
}

func NewPostgres(db *sqlx.DB) *Postgres {
	return &Postgres{
		db: db,
	}
}

func (p Postgres) PerformQuery(ammo int, ctx context.Context) error {
	myPerson := person.NewPerson()
	query := "SELECT name, dob, occupation, salary, address, city FROM persons WHERE id = $1;"
	err := p.db.GetContext(ctx, &myPerson, query, ammo)
	return err
}
