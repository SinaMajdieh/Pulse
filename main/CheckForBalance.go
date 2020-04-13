package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

func CheckForBalanceHandler(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles(CheckForBalanceAddress)
	temp.Execute(w, nil)
}

func BalanceResultHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		http.Redirect(w, r, "/Bank", http.StatusSeeOther)
	} else {
		CB := CheckBalance{}
		CB.UnmarshalData(r.Form)
		if CB.IsValid() == false {
			fmt.Fprintln(w, "Invalid Card Number Or Password")
		} else {
			temp, _ := template.ParseFiles(BalanceAddress)
			temp.Execute(w, CB)
		}
	}
}

func (CB *CheckBalance) UnmarshalData(form url.Values) {
	if x := form["cardnum"]; len(x) != 0 {
		CB.CardNum = x[0]
	}
	if x := form["password"]; len(x) != 0 {
		CB.Password = x[0]
	}
}

func (CB *CheckBalance) IsValid() bool {
	for _, v := range cards {
		if v.Cardnum == CB.CardNum && v.Password == CB.Password {
			CB.Name = v.Owner_name
			CB.LastName = v.Owner_lastname
			CB.Balance = v.Balance
			CB.BalanceString = v.Balance_string
			return true
		}
	}
	return false
}
