// Package models
// Time    : 2022/1/7 3:26 PM
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package models

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var schema = `
create table person (
    id int unsigned auto_increment primary key,
	first_name varchar(255),
	last_name varchar(255),
	email varchar(255)
);
`
var schema1 = `
CREATE TABLE place (
	id int unsigned auto_increment primary key,
    country varchar(255),
    city varchar(255),
    tel_code int
);`

type Person struct {
	Id        uint   `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
}

type Place struct {
	Id      uint           `db:"id"`
	Country string         `db:"country"`
	City    sql.NullString `db:"city"`
	TelCode int            `db:"tel_code"`
}

func MakeDBRequest() error {
	cfg := &mysql.Config{
		User:                 "mac",
		Passwd:               "workAholic!4",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "blog",
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	db, err := sqlx.Connect("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalln(err)
	}
	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv
	db.MustExec(schema)
	// db.MustExec(schema1)

	// tx := db.MustBegin()
	// tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES (?, ?, ?)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	// tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES (?, ?, ?)", "John", "Doe", "johndoeDNE@gmail.net")
	// tx.MustExec("INSERT INTO place (country, city, tel_code) VALUES (?, ?, ?)", "United States", "New York", "1")
	// tx.MustExec("INSERT INTO place (country, tel_code) VALUES (?, ?)", "Hong Kong", "852")
	// tx.MustExec("INSERT INTO place (country, tel_code) VALUES (?, ?)", "Singapore", "65")
	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	// tx.NamedExec("INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &Person{FirstName: "Jane", LastName: "Citizen", Email: "jane.citzen@example.com"})
	// tx.Commit()

	// Query the database, storing results in a []Person (wrapped in []interface{})
	people := []Person{}
	db.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")
	jason, john := people[0], people[1]

	fmt.Printf("%v\n%v\n", jason, john)

	// You can also get a single result, a la QueryRow
	jason = Person{}
	err = db.Get(&jason, "SELECT * FROM person WHERE first_name=?", "Jason")
	fmt.Printf("%#v\n", jason)

	// if you have null fields and use SELECT *, you must use sql.Null* in your struct
	places := []Place{}
	err = db.Select(&places, "SELECT * FROM place ORDER BY tel_code ASC")
	if err != nil {
		return err
	}
	usa, singsing, honkers := places[0], places[1], places[2]

	fmt.Printf("%#v\n%#v\n%#v\n", usa, singsing, honkers)

	// Loop through rows using only one struct
	place := Place{}
	rows, err := db.Queryx("SELECT * FROM place")
	for rows.Next() {
		err := rows.StructScan(&place)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", place)
	}

	// Named queries, using `:name` as the bindvar.  Automatic bindvar support
	// which takes into account the dbtype based on the driverName on sqlx.Open/Connect
	_, err = db.NamedExec(`INSERT INTO person (first_name,last_name,email) VALUES (:first,:last,:email)`,
		map[string]interface{}{
			"first": "Bin",
			"last":  "Smuth",
			"email": "bensmith@allblacks.nz",
		})

	return nil
}
