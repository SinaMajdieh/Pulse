package main

import (
	"net/http"
)

func main() {
	//var people People
	//var cards Cards
	ReloadDatas()
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/SearchResult", SearchHandler)
	http.HandleFunc("/DeepSearch", ShowProcessedData)
	http.HandleFunc("/Add", AddNewDataHandler)
	http.HandleFunc("/AddResult", AddingResult)
	http.HandleFunc("/Reload", ReloadDatasHandler)
	//Banks
	http.HandleFunc("/Bank", HandleBank)
	http.HandleFunc("/Bank/CheckForBalance", CheckForBalanceHandler)
	http.HandleFunc("/Bank/Balance", BalanceResultHandler)
	http.HandleFunc("/Bank/CreateAccount", HandleBankCreateAccountPage)
	http.HandleFunc("/Bank/CreateAccount/Result", HandleBankCreateAccountResultPage)
	http.HandleFunc("/Bank/CTC", HandleBankCTC)
	http.HandleFunc("/Bank/CTC/Ask", HandleBankCTCAsk)
	http.HandleFunc("/Bank/CTC/Result", HandleBankCTCResult)
	http.HandleFunc("/Bank/Deposit", HandleDeposit)
	http.HandleFunc("/Bank/Deposit/Ask", HandleDepositAsk)
	http.HandleFunc("/Bank/Deposit/Result", HandleDepositResult)
	http.HandleFunc("/Bank/Activity", LoadActivity)
	http.HandleFunc("/Bank/ActivityLogin", Activity)
	http.ListenAndServe(":8082", nil)
	//Find(people)
	//DonateProject(people , cards)
	//Hack(people , cards)
}
