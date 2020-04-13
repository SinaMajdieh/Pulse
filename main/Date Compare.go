package main

import (
	"strconv"
	"strings"
)

func (d *date)unmarshal(formattedDate string){
	//Format of date
	//01-02-2006 11:25:39 Mon
	falteredDate := strings.Split(formattedDate, " ")
	//Year Month Day
	YMD := strings.Split(falteredDate[0] , "-")
	//Hour Minute Second
	HMS := strings.Split(falteredDate[1] , ":")
	d[0] , _ = strconv.Atoi(YMD[2])
	d[1] , _ = strconv.Atoi(YMD[0])
	d[2] , _ = strconv.Atoi(YMD[1])

	d[3] , _ = strconv.Atoi(HMS[0])
	d[4] , _ = strconv.Atoi(HMS[1])
	d[5] , _ = strconv.Atoi(HMS[2])
}
func (d *Activities)swap(i int, j int) {
	x := (*d)[i]
	(*d)[i] = (*d)[j]
	(*d)[j] = x
}
func (d Activities) partitionDate(low int, high int) (partition int) {
	partition = low
	great := d[high]
	for i := low; i < high; i++ {

		if !d[i].IsBefore(great){
			d.swap(i, partition)
			partition++
		}
	}

	d.swap(partition, high)
	return
}
func (d Activities) sortDate(low int, high int) {
	if low < high {
		p := d.partitionDate(low, high)
		d.sortDate(low, p-1)
		d.sortDate(p+1, high)
	}
}
func (d Activities) SortByDate(){
	d.sortDate(0, len(d)-1)
}
