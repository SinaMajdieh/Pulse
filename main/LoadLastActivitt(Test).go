package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"

	digitsep "github.com/SinaMajdieh/DigitSeperator"
)

func Activity(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles(ActivityLoginAddress)
	temp.Execute(w, nil)
}
func LoadActivity(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		http.Redirect(w, r, "/Bank", http.StatusSeeOther)
	} else {
		login := Login{}
		login.unmarshal(r.Form)
		if login.IsValid() == false {
			fmt.Fprintln(w, "Invalid Card Number Or Password")
		} else {
			var activities Activities
			CTCactivities := LoadCTCRecords(login.CardNum)
			for _, v := range *CTCactivities {
				activities = append(activities, v)
			}
			DepositActivities := LoadDeposits(login.CardNum)
			for _, v := range *DepositActivities {
				activities = append(activities, v)
			}
			activities.SortByDate()
			var texts []string
			for i := 0; i < 10 && i < len(activities); i++ {
				if x, ok := activities[i].(CTCRecord); ok == true {
					text := ""
					if x.SenderCardNum == login.CardNum {
						text = fmt.Sprintf("%d.You Send $%s to %s %s (%s) on %s\n", i+1, x.MoneyString, x.ReceiverName, x.ReceiverLastName, x.ReceiverCardNum, x.DateString)
					} else if x.ReceiverCardNum == login.CardNum {
						text = fmt.Sprintf("%d.You Received $%s from %s %s (%s) on %s\n", i+1, x.MoneyString, x.SenderName, x.SenderLastName, x.SenderCardNum, x.DateString)
					}
					texts = append(texts, text)
				} else if x, ok := activities[i].(DepositRecord); ok == true {
					text := fmt.Sprintf("%d.You Deposited $%s on %s", i+1, x.MoneyString, x.DateString)
					texts = append(texts, text)
				}
			}
			temp, _ := template.ParseFiles(ActivityAddress)
			temp.Execute(w, texts)
		}
	}

}

func (l *Login) unmarshal(form url.Values) {
	if x := form["cardnum"]; len(x) != 0 {
		l.CardNum = x[0]
	}
	if x := form["password"]; len(x) != 0 {
		l.Password = x[0]
	}
}
func (l Login) IsValid() bool {
	for _, v := range cards {
		if v.Cardnum == l.CardNum && v.Password == l.Password {
			return true
		}
	}
	return false
}

func (ctc CTCRecord) IsBefore(activity1 activity) bool {
	for i, v := range ctc.Date {
		if v < activity1.ActivityDate()[i] {
			return true
		} else if v > activity1.ActivityDate()[i] {
			return false
		}
	}
	return false
}
func (ctc CTCRecord) ActivityDate() date {
	return ctc.Date
}

func (deposit DepositRecord) IsBefore(activity1 activity) bool {
	for i, v := range deposit.Date {
		if v < activity1.ActivityDate()[i] {
			return true
		} else if v > activity1.ActivityDate()[i] {
			return false
		}
	}
	return false
}
func (deposit DepositRecord) ActivityDate() date {
	return deposit.Date
}
func LoadCTCRecords(cardnum string) *[]CTCRecord {
	ctcRecords := make([]CTCRecord, 0)
	database := Open(BankRecordsDataBase)
	query := "SELECT * FROM " + BankRecordsTable + " WHERE sender_cardnum = ? OR receiver_cardnum = ?"
	rows, err := database.Query(query, cardnum, cardnum)
	if err != nil {
		fmt.Println("Error on reading " + BankRecordsTable + "table")
	} else {
		ctcRecord := CTCRecord{}
		for rows.Next() {
			err = rows.Scan(&ctcRecord.ID, &ctcRecord.SenderName, &ctcRecord.SenderLastName, &ctcRecord.SenderCardNum, &ctcRecord.ReceiverName, &ctcRecord.ReceiverLastName, &ctcRecord.ReceiverCardNum, &ctcRecord.Money, &ctcRecord.DateString)
			if err != nil {
				fmt.Println("Error on reading rows")
			} else {
				ctcRecord.MoneyString = digitsep.SepDigits(fmt.Sprintf("%.2f", ctcRecord.Money))
				ctcRecord.Date.unmarshal(ctcRecord.DateString)
				ctcRecords = append(ctcRecords, ctcRecord)
			}
		}

	}
	rows.Close()
	database.Close()
	return &ctcRecords
}
func LoadDeposits(cardnum string) *[]DepositRecord {
	depositRecords := make([]DepositRecord, 0)
	database := Open(BankRecordsDataBase)
	query := "SELECT * FROM " + BankDepositTable + " WHERE card_num = ?"
	rows, err := database.Query(query, cardnum)
	if err != nil {
		fmt.Println("Error on reading " + BankDepositTable + "table")
	} else {
		depositRecord := DepositRecord{}
		for rows.Next() {
			err = rows.Scan(&depositRecord.ID, &depositRecord.Name, &depositRecord.LastName, &depositRecord.CardNum, &depositRecord.Money, &depositRecord.DateString)
			if err != nil {
				fmt.Println("Error on reading rows")
			} else {
				depositRecord.MoneyString = digitsep.SepDigits(fmt.Sprintf("%.2f", depositRecord.Money))
				depositRecord.Date.unmarshal(depositRecord.DateString)
				depositRecords = append(depositRecords, depositRecord)
			}
		}
	}
	rows.Close()
	database.Close()
	return &depositRecords
}
