// Package basics
// Time    : 2022/8/7 17:18
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package basics

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Country string
	City    sql.NullString
	TelCode int
}

func GetSqliteConn(dbFile string) (db *sqlx.DB, err error) {
	db, err = sqlx.Open("sqlite3", dbFile)
	if err != nil {
		return nil, errors.Wrap(err, "read failed")
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "ping failed")
	}

	return db, nil
}

func CreateTables(db *sqlx.DB, schemes string) error {
	_, err := db.Exec(schemes)
	if err != nil {
		return err
	}
	return nil
}

func InsertPerson(db *sqlx.DB, persons []*Person) error {
	tx := db.MustBegin()
	for _, p := range persons {
		_, err := tx.NamedExec(`INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)`, p)
		if err != nil {
			err = tx.Rollback()
			return err
		}
	}
	_ = tx.Commit()
	return nil
}

func InsertPlace(db *sqlx.DB, places []*Place) error {
	tx := db.MustBegin()
	for _, p := range places {
		_, err := tx.Exec(`INSERT INTO place (country, city, telcode) VALUES ($1, $2, $3)`, p.Country, p.City, p.TelCode)
		if err != nil {
			err = tx.Rollback()
			return err
		}
	}
	_ = tx.Commit()
	return nil
}

func QueryPerson(db *sqlx.DB) ([]*Person, error) {
	dbFile := "./sqlx.db"
	db, err := GetSqliteConn(dbFile)
	if err != nil {
		return nil, err
	}
	var persons []*Person
	err = db.Select(&persons, "SELECT * FROM person ORDER BY first_name ASC")
	if err != nil {
		return nil, err
	}
	return persons, nil
}

func GetPerson(db *sqlx.DB, firstName string) (*Person, error) {
	var person Person
	err := db.Get(&person, "SELECT * FROM person WHERE first_name = $1", firstName)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func QueryXPlace(db *sqlx.DB) ([]*Place, error) {
	var places []*Place
	rows, err := db.Queryx("SELECT * FROM place")
	for rows.Next() {
		place := Place{}
		err = rows.StructScan(&place)
		if err != nil {
			return nil, err
		}
		places = append(places, &place)
	}
	return places, nil
}

func NamedQueryPerson(db *sqlx.DB) ([]*Person, error) {
	var persons []*Person
	rows, err := db.NamedQuery(`SELECT * FROM person WHERE first_name=:fn`, map[string]interface{}{"fn": "a"})
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		person := Person{}
		err = rows.StructScan(&person)
		if err != nil {
			return nil, err
		}
		persons = append(persons, &person)
	}
	return persons, nil
}
