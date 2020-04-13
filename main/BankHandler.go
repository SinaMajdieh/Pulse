package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	digitsep "github.com/SinaMajdieh/DigitSeperator"
)

func HandleBank(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles(BankAddress)
	temp.Execute(w, nil)
}

func HandleBankCreateAccountPage(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles(BankCreateAccountAddress)
	temp.Execute(w, nil)
}
func HandleBankCreateAccountResultPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	form := r.Form
	if r.Method == "GET" {
		http.Redirect(w, r, "/Bank", http.StatusSeeOther)
	} else {
		card := Card{}
		if x := form["FirstName"]; len(x) != 0 {
			card.Owner_name = CapitalizeFirstLetter(x[0])
		}
		if x := form["LastName"]; len(x) != 0 {
			card.Owner_lastname = CapitalizeFirstLetter(x[0])
		}
		if x := form["Password"]; len(x) != 0 {
			card.Password = x[0]
		}
		if x := form["Balance"]; len(x) != 0 {
			card.Balance, _ = strconv.ParseFloat(x[0], 64)
		}

		if found, hasaccount, record := SearchOnDataBase(card.Owner_name, card.Owner_lastname); found == true {
			cardnum := CreateCardNumber()
			card.Cardnum = cardnum
			AddToBank(card, hasaccount, record)
			fmt.Fprintf(w, ""+
				"Congratulation\n"+
				"Owner Name : %s\n"+
				"Owner Last Name : %s\n"+
				"Card Number : %s\n"+
				"Card Password : %s\n"+
				"Balanse : $%s\n",
				card.Owner_name,
				card.Owner_lastname,
				card.Cardnum,
				card.Password,
				digitsep.SepDigits(fmt.Sprintf("%.2f", card.Balance)),
			)
			ReloadDatas()
		} else {
			fmt.Fprintf(w, "%s %s NOT FOUND", card.Owner_name, card.Owner_lastname)
		}

	}
}

func SearchOnDataBase(name string, lastname string) (bool, bool, Record) {
	ReloadDatas()
	record := LoadRecord(CapitalizeFirstLetter(name), CapitalizeFirstLetter(lastname))
	name = strings.ToLower(name)
	lastname = strings.ToLower(lastname)
	FOUND := false
	HasAccount := false
	for _, v := range people {
		if strings.ToLower(v.Name) == name && strings.ToLower(v.LastName) == lastname {
			FOUND = true
			if v.Account.Cardnum != "-" {
				HasAccount = true
			}
			break
		}
	}
	return FOUND, HasAccount, record
}
func CreateCardNumber() string {
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)
	cardnum := ""
	for i := 0; i < 7; i++ {
		x := rand.Intn(10)
		cardnum += fmt.Sprint(x)
	}
	NEW := true
	for _, v := range cards {
		if cardnum == v.Cardnum {
			NEW = false
			break
		}
	}
	if NEW == false {
		cardnum = CreateCardNumber()
	}
	return cardnum
}
func CapitalizeFirstLetter(s string) string {
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}

func HandleBankCTC(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles(CTCAddress)
	temp.Execute(w, nil)
}

func HandleBankCTCResult(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/Bank/CTC", http.StatusSeeOther)
	}
	DoCTC(&GlobalCTC)
	temp, _ := template.ParseFiles("Templates/CTCDone.html")
	temp.Execute(w, nil)
	ReloadDatas()
}
func HandleBankCTCAsk(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		http.Redirect(w, r, "/Bank", http.StatusSeeOther)
	}
	form := r.Form
	ctc := CTC{}
	if x := form["CardNum"]; len(x) != 0 {
		ctc.CardNum = x[0]
	}
	if x := form["Password"]; len(x) != 0 {
		ctc.Password = x[0]
	}
	if x := form["ReceiverCardNum"]; len(x) != 0 {
		ctc.ReceiverCardNum = x[0]
	}
	if x := form["Money"]; len(x) != 0 {
		ctc.Money, _ = strconv.ParseFloat(x[0], 64)
		ctc.MoneyString = digitsep.SepDigits(x[0])
	}
	ValidCardNumAndPassword := CheckForCardNumAndPassword(&ctc)
	if ValidCardNumAndPassword == false {
		fmt.Fprintf(w, "Invalid Card Number Or Password")
	} else if CheckForCardNum(&ctc) == false {
		fmt.Fprintf(w, "Invalid Card Number")
	} else if CheckForMoney(&ctc) == false {
		fmt.Fprintf(w, "Sorry , You don't have enough money in your account")
	} else {
		GetTotalBalances(&ctc)
		GlobalCTC = ctc
		temp, _ := template.ParseFiles(CTCAskAddress)
		temp.Execute(w, ctc)
	}

}

func CheckForCardNumAndPassword(ctc *CTC) bool {
	CardNumValid := false
	PasswordMatchesCardNum := false
	for _, v := range cards {
		if v.Cardnum == ctc.CardNum {
			CardNumValid = true
			if v.Password == ctc.Password {
				PasswordMatchesCardNum = true
			}
			ctc.Balance = v.Balance
			ctc.Name = v.Owner_name
			ctc.LastName = v.Owner_lastname
			break
		}
	}
	return CardNumValid && PasswordMatchesCardNum
}
func CheckForCardNum(ctc *CTC) bool {
	CardNumValid := false
	for _, v := range cards {
		if v.Cardnum == ctc.ReceiverCardNum {
			CardNumValid = true
			ctc.ReceiverName = v.Owner_name
			ctc.ReceiverLastName = v.Owner_lastname
			ctc.ReceiverBalance = v.Balance
			ctc.ReceiverPassword = v.Password
			break
		}
	}
	return CardNumValid
}
func CheckForMoney(ctc *CTC) bool {
	MoneyIsValid := false
	if ctc.Balance >= ctc.Money {
		MoneyIsValid = true
	}
	return MoneyIsValid
}
func GetTotalBalances(ctc *CTC) {
	database := Open(BankDataBase)
	query := "SELECT balance FROM " + "records" + " WHERE name = ? AND lastname = ?"
	row := database.QueryRow(query, ctc.Name, ctc.LastName)
	_ = row.Scan(&ctc.TotalBalance)
	query = "SELECT balance FROM " + "records" + " WHERE name = ? AND lastname = ?"
	row = database.QueryRow(query, ctc.ReceiverName, ctc.ReceiverLastName)
	_ = row.Scan(&ctc.ReceiverTotalBalance)
	database.Close()
}
func DoCTC(ctc *CTC) {
	database := Open(BankRecordsDataBase)
	query := "INSERT INTO " + BankRecordsTable + " (sender_name, sender_last_name , sender_cardnum , receiver_name , receiver_last_name , receiver_cardnum , money , date) VALUES (? , ? , ? , ? , ? , ? , ? , ?)"
	tx, _ := database.Begin()
	dt := time.Now()
	date := fmt.Sprint(dt.Format("01-02-2006 15:04:05 Mon"))
	_, _ = tx.Exec(query, ctc.Name, ctc.LastName, ctc.CardNum, ctc.ReceiverName, ctc.ReceiverLastName, ctc.ReceiverCardNum, ctc.Money, date)
	tx.Commit()
	database.Close()

	database = Open(BankDataBase)
	query = "UPDATE " + "cards" + " SET balance = ? WHERE cardnumber = ? AND password = ? "
	tx, _ = database.Begin()
	_, _ = tx.Exec(query, ctc.Balance-ctc.Money, ctc.CardNum, ctc.Password)
	query = "UPDATE " + "cards" + " SET balance = ? WHERE cardnumber = ? AND password = ? "
	_, _ = tx.Exec(query, ctc.ReceiverBalance+ctc.Money, ctc.ReceiverCardNum, ctc.ReceiverPassword)

	query = "UPDATE " + "records" + " SET balance = ? WHERE name = ? AND lastname= ?"
	_, _ = tx.Exec(query, ctc.TotalBalance-ctc.Money, ctc.Name, ctc.LastName)
	query = "UPDATE " + "records" + " SET balance = ? WHERE name = ? AND lastname= ?"
	_, _ = tx.Exec(query, ctc.ReceiverTotalBalance+ctc.Money, ctc.ReceiverName, ctc.ReceiverLastName)
	tx.Commit()
	database.Close()

}
