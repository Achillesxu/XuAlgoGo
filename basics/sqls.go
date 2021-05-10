// Package basics
// Time    : 2021/5/7 8:57 下午
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package basics

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"time"
)

// Customer Class
type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}

// GetConnection method which returns sql.DB
func GetConnection() (database *sql.DB, err error) {
	databaseDriver := "mysql"
	databaseUser := "mac"
	databasePass := "tesT@2021"
	databaseName := "Golang"
	database, err = sql.Open(databaseDriver,
		databaseUser+":"+databasePass+"@/"+databaseName)
	if err != nil {
		return nil, errors.Wrap(err, "read failed")
	}
	// See "Important settings" section.
	database.SetConnMaxLifetime(time.Minute * 3)
	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(10)
	return database, nil
}

// GetCustomers method returns Customer Array
func GetCustomers() ([]Customer, error) {
	var database *sql.DB
	var err error
	var rows *sql.Rows
	database, err = GetConnection()
	if err != nil {
		return nil, err
	}

	rows, err = database.Query("SELECT * FROM Customer ORDER BY Customerid DESC")
	if err != nil {
		return nil, errors.Wrap(err, "GetCustomers failed")
	}
	var customer Customer
	customer = Customer{}
	var customers []Customer
	customers = []Customer{}
	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		err = rows.Scan(&customerId, &customerName, &ssn)
		if err != nil {
			panic(err.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers, customer)
	}
	return customers, nil
}

func InsertCustomer(customer Customer) (sql.Result, error) {
	var database *sql.DB
	var err error
	var insert *sql.Stmt

	database, err = GetConnection()
	if err != nil {
		return nil, err
	}

	insert, err = database.Prepare("INSERT INTO CUSTOMER(CustomerName, SSN) VALUES(?,?)")
	if err != nil {
		return nil, errors.Wrap(err, "InsertCustomer Prepare failed")
	}
	var res sql.Result
	res, err = insert.Exec(customer.CustomerName, customer.SSN)
	if err != nil {
		return nil, errors.Wrap(err, "InsertCustomer customer failed")
	}

	return res, nil
}

// UpdateCustomer method with parameter customer
func UpdateCustomer(customer Customer) error {
	var database *sql.DB
	var err error
	var update *sql.Stmt

	database, err = GetConnection()
	if err != nil {
		return err
	}

	update, err = database.Prepare("UPDATE CUSTOMER SET CustomerName=?, SSN=? WHERE CustomerId=?")
	if err != nil {
		return errors.Wrap(err, "UpdateCustomer customer failed")
	}
	_, err = update.Exec(customer.CustomerName, customer.SSN, customer.CustomerId)
	if err != nil {
		return errors.Wrap(err, "UpdateCustomer customer failed")
	}
	return nil
}

// DeleteCustomer method with parameter customer
func DeleteCustomer(customer Customer) error {
	var database *sql.DB
	var err error
	var del *sql.Stmt

	database, err = GetConnection()
	if err != nil {
		return err
	}
	del, err = database.Prepare("DELETE FROM Customer WHERE Customerid=?")
	if err != nil {
		return errors.Wrap(err, "DeleteCustomer customer failed")
	}
	_, err = del.Exec(customer.CustomerId)
	if err != nil {
		return err
	}
	return nil
}
