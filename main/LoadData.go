package main

import (
	"fmt"
	"strconv"

	digitsep "github.com/SinaMajdieh/DigitSeperator"
)

func (p People) Load(cards Cards) People {
	var people []person
	database := Open(CountryDataBase)
	query := "SELECT * FROM " + CityTable
	rows, err := database.Query(query)
	if nil != err {
		fmt.Println("Error on reading table : " + CityTable)
	}
	defer rows.Close()
	for rows.Next() {
		p := person{}
		err := rows.Scan(&p.Id, &p.Name, &p.LastName, &p.Birthday, &p.Mail, &p.Phone, &p.Gender, &p.Job, &p.JobRating)
		y := p.Birthday[0:4]
		year, _ := strconv.Atoi(y)
		p.Age = YEAR - year
		p.Account = cards.FindCard(p.Name, p.LastName)
		if err != nil {
			fmt.Println("Error on reading row")
		} else {
			people = append(people, p)
		}

	}
	database.Close()
	return people
}
func LoadRecord(name string, lastname string) Record {
	var record Record
	database := Open(BankDataBase)
	query := "SELECT * FROM " + "records" + " WHERE name = ? AND lastname = ?"
	row := database.QueryRow(query, name, lastname)
	_ = row.Scan(&record.Id, &record.name, &record.lastName, &record.balance, &record.accountNum)
	database.Close()
	return record
}

func LoadRecords() []Record {
	var records []Record
	database := Open(BankDataBase)
	query := "SELECT * From " + BankRecordsTable
	rows, err := database.Query(query)
	if err != nil {
		fmt.Println("Error on reading table : " + "records")
	}
	defer rows.Close()
	for rows.Next() {
		record := Record{}
		err := rows.Scan(&record.Id, &record.name, &record.lastName, &record.balance, &record.accountNum)
		if err != nil {
			fmt.Println("Error on reading row")
		} else {
			records = append(records, record)
		}
	}
	database.Close()
	return records
}

func (c Cards) Load() Cards {
	var cards []Card
	database := Open(BankDataBase)
	query := "SELECT Owner_name , Owner_lastname , cardnumber , Password , Balance FROM " + BankCardsTable
	rows, err := database.Query(query)
	if nil != err {
		fmt.Println("Error on reading table : " + "cards")
	}
	defer rows.Close()
	for rows.Next() {
		c := Card{}
		err := rows.Scan(&c.Owner_name, &c.Owner_lastname, &c.Cardnum, &c.Password, &c.Balance)
		c.Balance_string = digitsep.SepDigits(fmt.Sprintf("%.2f", c.Balance))
		if err != nil {
			fmt.Println("Error on reading row")
		} else {
			cards = append(cards, c)
		}

	}
	database.Close()
	return cards

}
