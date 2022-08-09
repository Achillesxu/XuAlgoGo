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
