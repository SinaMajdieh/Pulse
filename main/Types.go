package main

import (
	"fmt"
	"strconv"
)

type Cards []Card
type People []person
type Datas []FullData
type date [6]int
type Activities []activity

type Card struct {
	Owner_name     string
	Owner_lastname string
	Cardnum        string
	Password       string
	Balance        float64
	Balance_string string
}

type person struct {
	Id        int
	Name      string
	LastName  string
	Age       int
	Birthday  string
	Job       string
	JobRating int
	Mail      string
	Phone     string
	Gender    string
	Account   *Card
}

func (p person) Print() {
	fmt.Printf("%-5s : %-10s %-10s %-5s %-30s %-15s %-10s %-15s %-70s\n", strconv.Itoa(p.Id), p.Name, p.LastName, strconv.Itoa(p.Age), p.Mail, p.Phone, p.Gender, strconv.Itoa(p.JobRating), p.Job)
}

type FullData struct {
	Person person
	Card   Card
}

type activity interface {
	IsBefore(activity) bool
	ActivityDate() date
}

type Record struct {
	Id         int
	name       string
	lastName   string
	balance    float64
	accountNum int
}

type CTC struct {
	CardNum              string
	Password             string
	Name                 string
	LastName             string
	Balance              float64
	TotalBalance         float64
	ReceiverCardNum      string
	ReceiverPassword     string
	ReceiverName         string
	ReceiverLastName     string
	ReceiverBalance      float64
	ReceiverTotalBalance float64
	Money                float64
	MoneyString          string
}
type CTCRecord struct {
	ID               int
	SenderName       string
	SenderLastName   string
	SenderCardNum    string
	ReceiverCardNum  string
	ReceiverName     string
	ReceiverLastName string
	Money            float64
	MoneyString      string
	DateString       string
	Date             date
}

type CheckBalance struct {
	CardNum       string
	Password      string
	Name          string
	LastName      string
	Balance       float64
	BalanceString string
}

type Deposit struct {
	Name         string
	LastName     string
	CardNum      string
	Password     string
	Money        float64
	Date         string
	TotalBalance float64
	Balance      float64
}

type DepositRecord struct {
	ID          int
	Name        string
	LastName    string
	CardNum     string
	Money       float64
	MoneyString string
	Date        date
	DateString  string
}

type Result struct {
	People Datas
	Count  string
}

type Login struct {
	CardNum  string
	Password string
}
