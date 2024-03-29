// Package basics
// Time    : 2022/8/8 23:11
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package basics

import (
	"database/sql"
	"testing"
)

func TestGetSqliteConn(t *testing.T) {
	dbFile := "./sqlx.db"
	db, err := GetSqliteConn(dbFile)
	if err != nil {
		t.Error(err)
	}
	t.Log(db)
}

func TestCreateTables(t *testing.T) {
	dbFile := "./sqlx.db"
	db, err := GetSqliteConn(dbFile)
	if err != nil {
		t.Error(err)
	}
	schemes := `
DROP TABLE IF EXISTS person;
CREATE TABLE person
(
    first_name text,
    last_name  text,
    email      text
);

Drop Table if EXISTS place;
CREATE TABLE place
(
    country text,
    city    text NULL,
    telcode integer
);`
	err = CreateTables(db, schemes)
	if err != nil {
		t.Error(err)
	}
}

func TestInsertPerson(t *testing.T) {
	var testPersons = []*Person{
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i"},
	}
	dbFile := "./sqlx.db"
	db, err := GetSqliteConn(dbFile)
	if err != nil {
		t.Error(err)
	}
	err = InsertPerson(db, testPersons)
	if err != nil {
		t.Error(err)
	}
}

func TestInsertPlace(t *testing.T) {
	var testPlaces = []*Place{
		{Country: "a", City: sql.NullString{String: "b", Valid: true}, TelCode: 1},
		{Country: "d", City: sql.NullString{Valid: false}, TelCode: 2},
		{Country: "g", City: sql.NullString{Valid: true}, TelCode: 3},
	}
	dbFile := "./sqlx.db"
	db, err := GetSqliteConn(dbFile)
	if err != nil {
		t.Error(err)
	}
	err = InsertPlace(db, testPlaces)
	if err != nil {
		t.Error(err)
	}
}

func TestQueryPerson(t *testing.T) {
	dbFile := "./sqlx.db"
	db, err := GetSqliteConn(dbFile)
	if err != nil {
		t.Error(err)
	}
	persons, err := QueryPerson(db)
	if err != nil {
		t.Error(err)
	}
	for _, p := range persons {
		t.Logf("%#v", p)
	}
}

func TestGetPerson(t *testing.T) {
	dbFile := "./sqlx.db"
	db, err := GetSqliteConn(dbFile)
	if err != nil {
		t.Error(err)
	}
	person, err := GetPerson(db, "a")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v", person)
}

func TestQueryXPlace(t *testing.T) {
	dbFile := "./sqlx.db"
	db, err := GetSqliteConn(dbFile)
	if err != nil {
		t.Error(err)
	}
	places, err := QueryXPlace(db)
	if err != nil {
		t.Error(err)
	}
	for _, p := range places {
		t.Logf("%#v", p)
	}
}

func TestNamedQueryPerson(t *testing.T) {
	dbFile := "./sqlx.db"
	db, err := GetSqliteConn(dbFile)
	if err != nil {
		t.Error(err)
	}
	persons, err := NamedQueryPerson(db)
	if err != nil {
		t.Error(err)
	}
	for _, p := range persons {
		t.Logf("%#v", p)
	}
}
