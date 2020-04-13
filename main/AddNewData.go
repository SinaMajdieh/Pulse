package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)


func AddNewDataHandler(w http.ResponseWriter , r *http.Request){
	temp , _ := template.ParseFiles(AddNewDataAddress)
	temp.Execute(w , nil)
}
func AddingResult(w http.ResponseWriter , r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		http.Redirect(w , r ,"/" , http.StatusSeeOther)
	} else {
		p := person{}
		if x := r.Form["FirstName"]; len(x) != 0 {
			p.Name = CapitalizeFirstLetter(x[0])
		}
		if x := r.Form["LastName"]; len(x) != 0 {
			p.LastName = CapitalizeFirstLetter(x[0])
		}
		if x := r.Form["Prefix"]; len(x) != 0 {
			p.Phone = x[0]
		}
		if x := r.Form["Phone"]; len(x) != 0 {
			p.Phone += x[0]
		}
		if x := r.Form["Mail"]; len(x) != 0 {
			p.Mail = x[0]
		}
		if x := r.Form["BirthDate"]; len(x) != 0 {
			p.Birthday = x[0]
			Date := strings.Split(p.Birthday, "-")
			year, _ := strconv.Atoi(Date[0])
			p.Age = YEAR - year
		}
		if x := r.Form["Job"]; len(x) != 0 {
			p.Job = x[0]
		}
		if x := r.Form["JobRating"]; len(x) != 0 {
			p.JobRating, _ = strconv.Atoi(x[0])
		}
		if x := r.Form["Gender"]; len(x) != 0 {
			p.Gender = x[0]
		}
		AddToDataBase(p)
		fmt.Fprintf(w , "Data Was Successfuly Saved\n" +
			"Name : %s\n" +
			"Last Name : %s\n" +
			"Phone : %s\n" +
			"Mail : %s\n" +
			"Birth Date : %s\n" +
			"Job : %s\n" +
			"Job Rating : %d\n" +
			"Gender : %s\n",
			p.Name,
			p.LastName,
			p.Phone,
			p.Mail,
			p.Birthday,
			p.Job,
			p.JobRating,
			p.Gender,
		)

	}
	ReloadDatas()
}
