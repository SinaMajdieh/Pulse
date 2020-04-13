package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func HandleDeposit(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles(DepositAddress)
	temp.Execute(w, nil)
}

func HandleDepositResult(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/Bank/Deposit", http.StatusSeeOther)
	} else {
		GlobalDeposit.DepositMoney()
		temp, _ := template.ParseFiles(DepositResultAddress)
		temp.Execute(w, GlobalDeposit)
		ReloadDatas()
	}
}
func HandleDepositAsk(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/Bank/Deposit", http.StatusSeeOther)
	} else {
		_ = r.ParseForm()
		deposit := Deposit{}
		UnmarshalDepositForm(r.Form, &deposit)
		if deposit.IsValid() == false {
			fmt.Fprintf(w, "Invalid Card Number Or Password")
		} else {
			GlobalDeposit = &deposit
			temp, _ := template.ParseFiles(DepositAskAddress)
			temp.Execute(w, *GlobalDeposit)
		}
	}

}

func UnmarshalDepositForm(form url.Values, deposit *Deposit) {
	if x := form["cardnum"]; len(x) != 0 {
		deposit.CardNum = x[0]
	}
	if x := form["password"]; len(x) != 0 {
		deposit.Password = x[0]
	}
	if x := form["money"]; len(x) != 0 {
		deposit.Money, _ = strconv.ParseFloat(x[0], 64)
	}
	date := time.Now()
	deposit.Date = date.Format("01-02-2006 15:04:05 Mon")
}
func (d *Deposit) IsValid() bool {
	for _, v := range cards {
		if v.Cardnum == d.CardNum && v.Password == d.Password {
			d.Name = v.Owner_name
			d.LastName = v.Owner_lastname
			d.Balance = v.Balance
			d.TotalBalance = FindRecord(d.Name, d.LastName).balance
			return true
		}
	}
	return false
}

func (d Deposit) DepositMoney() {
	database := Open(BankRecordsDataBase)
	query := "INSERT INTO " + BankDepositTable + " (name, last_name , card_num , money , date) VALUES (? , ? , ? , ? , ?)"
	tx, _ := database.Begin()
	_, _ = tx.Exec(query, d.Name, d.LastName, d.CardNum, d.Money, d.Date)
	tx.Commit()
	database.Close()

	database = Open(BankDataBase)
	tx, _ = database.Begin()
	query = "UPDATE " + "cards" + " SET balance = ? WHERE cardnumber = ? AND password = ? "
	_, _ = tx.Exec(query, d.Balance+d.Money, d.CardNum, d.Password)
	query = "UPDATE " + "records" + " SET balance = ? WHERE name = ? AND lastname= ?"
	_, _ = tx.Exec(query, d.TotalBalance+d.Money, d.Name, d.LastName)
	tx.Commit()
	database.Close()

}

func FindRecord(name string, lastname string) *Record {
	for _, v := range records {
		if v.name == name && v.lastName == lastname {
			return &v
		}
	}
	return &Record{}
}
