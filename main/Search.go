package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	digitsep "github.com/SinaMajdieh/DigitSeperator"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("Templates/SearchPage.html")
	temp.Execute(w, nil)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		temp, _ := template.ParseFiles("Templates/SearchPage.html")
		temp.Execute(w, nil)
	} else {
		SFFN := Name{}
		if len(r.Form["FNV"]) != 0 {
			SFFN.Value = r.Form["FNV"][0]
		}
		if len(r.Form["SFFN"]) != 0 {
			SFFN.Status = r.Form["SFFN"][0]
		}
		if len(r.Form["SFFNops"]) != 0 {
			SFFN.Type = r.Form["SFFNops"][0]
		}
		SFLN := Name{}
		if len(r.Form["LNV"]) != 0 {
			SFLN.Value = r.Form["LNV"][0]
		}
		if len(r.Form["SFLN"]) != 0 {
			SFLN.Status = r.Form["SFLN"][0]
		}
		if len(r.Form["SFLNops"]) != 0 {
			SFLN.Type = r.Form["SFLNops"][0]
		}
		SFG := Gender{}
		if len(r.Form["SFG"]) != 0 {
			SFG.Status = r.Form["SFG"][0]
		}
		if len(r.Form["SFGops"]) != 0 {
			SFG.Type = r.Form["SFGops"][0]
		}
		SFMPN := MobilePhoneNumber{}
		if len(r.Form["SFMPN"]) != 0 {
			SFMPN.Status = r.Form["SFMPN"][0]
		}
		if x := r.Form["SFMPNops"]; len(x) != 0 {
			SFMPN.Type = x[0]
		}
		if x := r.Form["SFMPNpre"]; len(x) != 0 {
			SFMPN.Prefix = x[0]
		}
		if x := r.Form["MPNV"]; len(x) != 0 {
			SFMPN.Value = x[0]
		}
		SFEA := Name{}
		if x := r.Form["SFEA"]; len(x) != 0 {
			SFEA.Status = x[0]
		}
		if x := r.Form["SFEAops"]; len(x) != 0 {
			SFEA.Type = x[0]
		}
		if x := r.Form["EAV"]; len(x) != 0 {
			SFEA.Value = x[0]
		}
		SFJ := Name{}
		if x := r.Form["SFJ"]; len(x) != 0 {
			SFJ.Status = x[0]
		}
		if x := r.Form["SFJops"]; len(x) != 0 {
			SFJ.Type = x[0]
		}
		if x := r.Form["JV"]; len(x) != 0 {
			SFJ.Value = x[0]
		}
		SFJR := Number{}
		if x := r.Form["SFJR"]; len(x) != 0 {
			SFJR.Status = x[0]
		}
		if x := r.Form["SFJRops"]; len(x) != 0 {
			SFJR.Type = x[0]
		}
		if x := SFJR.Type; x == "equals" && len(r.Form["JRVEquals"]) != 0 {
			SFJR.Value, _ = strconv.Atoi(r.Form["JRVEquals"][0])
		} else if x == "from-to" && (len(r.Form["JRVFrom"]) != 0 && len(r.Form["AVTo"]) != 0) {
			SFJR.ValueFrom, _ = strconv.Atoi(r.Form["JRVFrom"][0])
			SFJR.ValueTo, _ = strconv.Atoi(r.Form["JRVTo"][0])
		}
		SFA := Number{}
		if x := r.Form["SFA"]; len(x) != 0 {
			SFA.Status = x[0]
		}
		if x := r.Form["SFAops"]; len(x) != 0 {
			SFA.Type = x[0]
		}
		if x := SFA.Type; x == "equals" && len(r.Form["AVEquals"]) != 0 {
			SFA.Value, _ = strconv.Atoi(r.Form["AVEquals"][0])
		} else if x == "from-to" && (len(r.Form["AVFrom"]) != 0 && len(r.Form["AVTo"]) != 0) {
			SFA.ValueFrom, _ = strconv.Atoi(r.Form["AVFrom"][0])
			SFA.ValueTo, _ = strconv.Atoi(r.Form["AVTo"][0])
		}
		SFI := Number{}
		if x := r.Form["SFI"]; len(x) != 0 {
			SFI.Status = x[0]
		}
		if x := r.Form["SFIops"]; len(x) != 0 {
			SFI.Type = x[0]
		}
		if x := SFI.Type; x == "equals" && len(r.Form["IVEquals"]) != 0 {
			SFI.Value, _ = strconv.Atoi(r.Form["IVEquals"][0])
		} else if x == "from-to" && (len(r.Form["IVFrom"]) != 0 && len(r.Form["IVTo"]) != 0) {
			SFI.ValueFrom, _ = strconv.Atoi(r.Form["IVFrom"][0])
			SFI.ValueTo, _ = strconv.Atoi(r.Form["IVTo"][0])
		}
		SFB := Float{}
		if x := r.Form["SFB"]; len(x) != 0 {
			SFB.Status = x[0]
		}
		if x := r.Form["SFBops"]; len(x) != 0 {
			SFB.Type = x[0]
		}
		if x := SFB.Type; x == "equals" && len(r.Form["BVEquals"]) != 0 {
			SFB.Value, _ = strconv.ParseFloat(r.Form["BVEquals"][0], 64)
		} else if x == "from-to" && (len(r.Form["BVFrom"]) != 0 && len(r.Form["BVTo"]) != 0) {
			SFB.ValueFrom, _ = strconv.ParseFloat(r.Form["BVFrom"][0], 64)
			SFB.ValueTo, _ = strconv.ParseFloat(r.Form["BVTo"][0], 64)
		}
		SFCN := Name{}
		if x := r.Form["SFCN"]; len(x) != 0 {
			SFCN.Status = x[0]
		}
		if x := r.Form["SFCNops"]; len(x) != 0 {
			SFCN.Type = x[0]
		}
		if x := r.Form["CNV"]; len(x) != 0 {
			SFCN.Value = x[0]
		}
		SFCP := Name{}
		if x := r.Form["SFCP"]; len(x) != 0 {
			SFCP.Status = x[0]
		}
		if x := r.Form["SFCPops"]; len(x) != 0 {
			SFCP.Type = x[0]
		}
		if x := r.Form["CPV"]; len(x) != 0 {
			SFCP.Value = x[0]
		}
		SA := SearchAttributes{
			SFFN,
			SFLN,
			SFG,
			SFMPN,
			SFEA,
			SFJ,
			SFJR,
			SFA,
			SFI,
			SFB,
			SFCN,
			SFCP,
		}
		SearchedPeople, count := SearchDataBase( /*people , cards ,*/ SA)
		result := Result{
			People: SearchedPeople,
			Count:  digitsep.SepDigits(fmt.Sprint(count)),
		}
		if len(r.Form["sort_by_balance"]) != 0 {
			result.People.SortByBalance()
		} else if len(r.Form["sort_by_age"]) != 0 {
			result.People.SortByAge()
		} else if len(r.Form["sort_by_job_rating"]) != 0 {
			result.People.SortByJobRating()
		} else if len(r.Form["sort_by_id"]) != 0 {
			result.People.SortByID()
		} else if len(r.Form["sort_by_name"]) != 0 {
			result.People.SortByName()
		} else if len(r.Form["sort_by_last_name"]) != 0 {
			result.People.SortByLastName()
		}

		template, _ := template.ParseFiles("Templates/main.html")
		template.Execute(w, result)
	}
}
func ShowProcessedData(w http.ResponseWriter, r *http.Request) {
	SearchedPeople, count := Search( /*people , cards*/ )
	result := Result{
		People: SearchedPeople,
		Count:  digitsep.SepDigits(fmt.Sprint(count)),
	}
	//result.People.SortByBalance()
	template, _ := template.ParseFiles("Templates/main.html")
	template.Execute(w, result)
}

func Search( /*people People , cards Cards*/ ) (Datas, int) {
	var totalMoney float64 = 0
	var SearchedPeople Datas
	count := 0
	for _, v := range people {
		if v.Account.Password == "3299" {
			totalMoney += v.Account.Balance
			data := FullData{
				Person: v,
				Card:   *v.Account,
			}
			SearchedPeople = append(SearchedPeople, data)
			count++
		}
	}
	fmt.Printf("%.2f", totalMoney/float64((count)))
	return SearchedPeople, count
}
func SearchDataBase( /*people People  , cards Cards ,*/ SA SearchAttributes) (Datas, int) {
	var SearchedPeople Datas
	count := 0
	for _, v := range people {
		if SA.SFFN.Status == "on" {
			if SA.SFFN.Type == "contains" {
				if strings.Contains(strings.ToLower(v.Name), strings.ToLower(SA.SFFN.Value)) == false {
					continue
				}
			} else if SA.SFFN.Type == "equals" {
				if strings.ToLower(v.Name) != strings.ToLower(SA.SFFN.Value) {
					continue
				}
			}
		}
		if SA.SFLN.Status == "on" {
			if SA.SFLN.Type == "contains" {
				if strings.Contains(strings.ToLower(v.LastName), strings.ToLower(SA.SFLN.Value)) == false {
					continue
				}
			} else if SA.SFLN.Type == "equals" {
				if strings.ToLower(v.LastName) != strings.ToLower(SA.SFLN.Value) {
					continue
				}
			}
		}
		if SA.SFG.Status == "on" {
			if v.Gender != SA.SFG.Type {
				continue
			}
		}
		if SA.SFMPN.Status == "on" {
			if !strings.HasPrefix(v.Phone, SA.SFMPN.Prefix) {
				continue
			}
			if SA.SFMPN.Type == "contains" {
				if !strings.Contains(v.Phone, SA.SFMPN.Value) {
					continue
				}
			} else if SA.SFMPN.Type == "equals" {
				if v.Phone != SA.SFMPN.Prefix+SA.SFMPN.Value {
					continue
				}
			}
		}
		if SA.SFEA.Status == "on" {
			if t := SA.SFEA.Type; t == "contains" {
				if !strings.Contains(strings.ToLower(v.Mail), strings.ToLower(SA.SFEA.Value)) {
					continue
				}
			} else if t == "equals" {
				if strings.ToLower(v.Mail) != strings.ToLower(SA.SFEA.Value) {
					continue
				}
			}
		}
		if SA.SFJ.Status == "on" {
			if t := SA.SFJ.Type; t == "contains" {
				if !strings.Contains(strings.ToLower(v.Job), strings.ToLower(SA.SFJ.Value)) {
					continue
				}
			} else if t == "equals" {
				if strings.ToLower(v.Job) != strings.ToLower(SA.SFJ.Value) {
					continue
				}
			}
		}
		if SA.SFJR.Status == "on" {
			if t := SA.SFJR.Type; t == "equals" {
				if v.JobRating != SA.SFJR.Value {
					continue
				}
			} else if t == "from-to" {
				if !(v.JobRating >= SA.SFJR.ValueFrom && v.JobRating <= SA.SFJR.ValueTo) {
					continue
				}
			}
		}
		if SA.SFA.Status == "on" {
			if t := SA.SFA.Type; t == "equals" {
				if v.Age != SA.SFA.Value {
					continue
				}
			} else if t == "from-to" {
				if !(v.Age >= SA.SFA.ValueFrom && v.Age <= SA.SFA.ValueTo) {
					continue
				}
			}
		}
		if SA.SFI.Status == "on" {
			if t := SA.SFI.Type; t == "equals" {
				if v.Id != SA.SFI.Value {
					continue
				}
			} else if t == "from-to" {
				if !(v.Id >= SA.SFI.ValueFrom && v.Id <= SA.SFI.ValueTo) {
					continue
				}
			}
		}
		if SA.SFB.Status == "on" {
			if t := SA.SFB.Type; t == "equals" {
				if v.Account.Balance != SA.SFB.Value {
					continue
				}
			} else if t == "from-to" {
				if !(v.Account.Balance >= SA.SFB.ValueFrom && v.Account.Balance <= SA.SFB.ValueTo) {
					continue
				}
			}
		}
		if SA.SFCN.Status == "on" {
			if t := SA.SFCN.Type; t == "contains" {
				if !strings.Contains(strings.ToLower(v.Account.Cardnum), strings.ToLower(SA.SFCN.Value)) {
					continue
				}
			} else if t == "equals" {
				if strings.ToLower(v.Account.Cardnum) != strings.ToLower(SA.SFCN.Value) {
					continue
				}
			}
		}
		if SA.SFCP.Status == "on" {
			if t := SA.SFCP.Type; t == "contains" {
				if !strings.Contains(strings.ToLower(v.Account.Password), strings.ToLower(SA.SFCP.Value)) {
					continue
				}
			} else if t == "equals" {
				if strings.ToLower(v.Account.Password) != strings.ToLower(SA.SFCP.Value) {
					continue
				}
			}
		}
		data := FullData{
			Person: v,
			Card:   *(v.Account),
		}
		SearchedPeople = append(SearchedPeople, data)
		count++
	}
	return SearchedPeople, count
}

func Find(people []person) {
	fmt.Printf("%-5s : %-10s %-10s %-5s %-30s %-15s %-10s %-15s %-70s\n", "ID", "NAME", "LAST NAME", "AGE", "MAIL", "PHONE", "GENDER", "JOB RATING", "JOB")
	count := 0
	specialCount := 0
	for _, v := range people {
		if v.Age >= 30 && v.Age <= 35 && v.Gender == "male" && v.JobRating >= 5 && v.JobRating <= 7 && v.LastName != "Hill" && v.LastName != "Morris" && strings.Contains(strings.ToLower(v.Job), "food") {
			if v.JobRating >= 5 {
				v.Print()
				specialCount++
				continue
			}
			v.Print()
			count++
		}
	}
	fmt.Println(strconv.Itoa(count) + " result(s)")
	fmt.Println(strconv.Itoa(specialCount) + " special result(s)")
}

func ReloadDatas() {
	cards = cards.Load()
	people = people.Load(cards)
	records = LoadRecords()
}
func ReloadDatasHandler(w http.ResponseWriter, r *http.Request) {
	ReloadDatas()
	temp, _ := template.ParseFiles("Templates/SearchPage.html")
	temp.Execute(w, nil)
}

func (c Cards) FindCard(name string, lastname string) *Card {
	for _, v := range c {
		if v.Owner_name == name && v.Owner_lastname == lastname {
			return &v
		}
	}

	emptyCard := Card{
		Owner_name:     name,
		Owner_lastname: lastname,
		Cardnum:        "-",
		Password:       "-",
		Balance:        0,
		Balance_string: "0",
	}

	return &emptyCard
}
