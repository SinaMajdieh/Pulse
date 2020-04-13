package main

const (
	YEAR = 2020

	//Templates
	AddNewDataAddress        = "Templates/AddNewDataPage.html"
	BankCreateAccountAddress = "Templates/BankCreateAccountPage.html"
	CTCAddress               = "Templates/CTC.html"
	CTCAskAddress            = "Templates/CTCAsk.html"
	CTCDoneAddress           = "Templates/CTCDone.html"
	DepositAddress           = "Templates/DepositPage.html"
	DepositAskAddress        = "Templates/DepositAskPage.html"
	DepositResultAddress     = "Templates/DepositResultPage.html"
	SearchResultAddress      = "Templates/main.html"
	SearchAddress            = "Templates/SearchPage.html"
	BankAddress              = "Templates/Bank.html"
	CheckForBalanceAddress   = "Templates/CheckForBalance.html"
	BalanceAddress           = "Templates/Balance.html"
	ActivityLoginAddress     = "Templates/ActivityLogin.html"
	ActivityAddress          = "Templates/Activity.html"

	//DataBases
	BankRecordsDataBase = "DataBases/SinaBankRecords.sqlite"
	CountryDataBase     = "DataBases/United States of America.sqlite"
	BankDataBase        = "DataBases/Sina.sqlite"

	//DataBase Tables
	BankRecordsTable = "records"
	BankCardsTable   = "cards"
	BankDepositTable = "depositOrwithdraw"
	CityTable        = "New_York"
)

var (
	GlobalCTC     CTC
	GlobalDeposit *Deposit
	people        People
	cards         Cards
	records       []Record
)
