package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Open(name string) *sql.DB {
	database, err := sql.Open("sqlite3", name)
	if nil != err {
		fmt.Println("Error on opening database")
	}
	return database
}

func AddToDataBase(p person) {
	database := Open(CountryDataBase)
	query := "INSERT INTO " + CityTable + " (firstname, lastname , birthday , mail , phone , gender , job , job_rating) VALUES (? , ? , ? , ? , ? , ? , ? , ?)"
	tx, _ := database.Begin()
	_, _ = tx.Exec(query, p.Name, p.LastName, p.Birthday, p.Mail, p.Phone, p.Gender, p.Job, p.JobRating)
	tx.Commit()
	database.Close()
}

func AddToBank(c Card, hasaccount bool, record Record) {
	database := Open(BankDataBase)
	query := "INSERT INTO " + "cards" + " (owner_name, owner_lastname , cardnumber , password , balance) VALUES (? , ? , ? , ? , ? )"
	tx, _ := database.Begin()
	_, _ = tx.Exec(query, c.Owner_name, c.Owner_lastname, c.Cardnum, c.Password, c.Balance)
	if hasaccount == true {

		query = "UPDATE " + "records" + " SET accountnum=? , balance=? WHERE name=? AND lastname=?"
		_, _ = tx.Exec(query, record.accountNum+1, record.balance+c.Balance, record.name, record.lastName)
	} else {
		query := "INSERT INTO " + "records" + " (name, lastname , balance , accountnum) VALUES (? , ? , ? , ?)"
		_, _ = tx.Exec(query, c.Owner_name, c.Owner_lastname, c.Balance, 1)
	}
	tx.Commit()
	database.Close()
}
