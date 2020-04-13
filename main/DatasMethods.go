package main

func (d *Datas)swap(i int, j int) {
	x := (*d)[i]
	(*d)[i] = (*d)[j]
	(*d)[j] = x
}
func (d Datas) partitionBalance(low int, high int) (partition int) {
	partition = low
	great := d[high]
	for i := low; i < high; i++ {

		if d[i].Card.Balance < great.Card.Balance {
			d.swap(i, partition)
			partition++
		}else if d[i].Card.Balance == great.Card.Balance{
			if d[i].Person.Age > great.Person.Age{
				d.swap(i,partition)
				partition++
			}
		}
	}

	d.swap(partition, high)
	return
}

func (d Datas) sortBalance(low int, high int) {
	if low < high {
		p := d.partitionBalance(low, high)
		d.sortBalance(low, p-1)
		d.sortBalance(p+1, high)
	}
}
func (d Datas) SortByBalance(){
	d.sortBalance(0, len(d)-1)
}
func (d Datas) partitionAge(low int, high int) (partition int) {
	partition = low
	great := d[high]
	for i := low; i < high; i++ {

		if d[i].Person.Age < great.Person.Age {
			d.swap(i, partition)
			partition++
		}
	}

	d.swap(partition, high)
	return
}
func (d Datas) sortAge(low int, high int) {
	if low < high {
		p := d.partitionAge(low, high)
		d.sortAge(low, p-1)
		d.sortAge(p+1, high)
	}
}
func (d Datas) SortByAge(){
	d.sortAge(0, len(d)-1)
}

//Sort By Job Rating
func (d Datas) partitionJobRating(low int, high int) (partition int) {
	partition = low
	great := d[high]
	for i := low; i < high; i++ {

		if d[i].Person.JobRating < great.Person.JobRating {
			d.swap(i, partition)
			partition++
		}
	}

	d.swap(partition, high)
	return
}
func (d Datas) sortJobRating(low int, high int) {
	if low < high {
		p := d.partitionJobRating(low, high)
		d.sortJobRating(low, p-1)
		d.sortJobRating(p+1, high)
	}
}
func (d Datas) SortByJobRating(){
	d.sortJobRating(0, len(d)-1)
}

//Sort By ID
func (d Datas) partitionID(low int, high int) (partition int) {
	partition = low
	great := d[high]
	for i := low; i < high; i++ {

		if d[i].Person.Id < great.Person.Id {
			d.swap(i, partition)
			partition++
		}
	}

	d.swap(partition, high)
	return
}
func (d Datas) sortID(low int, high int) {
	if low < high {
		p := d.partitionID(low, high)
		d.sortID(low, p-1)
		d.sortID(p+1, high)
	}
}
func (d Datas) SortByID(){
	d.sortID(0, len(d)-1)
}

//Sort By Name
func (d Datas) partitionName(low int, high int) (partition int) {
	partition = low
	great := d[high]
	for i := low; i < high; i++ {

		if x := ComesBefore(d[i].Person.Name , great.Person.Name); x == "true" {
			d.swap(i, partition)
			partition++
		}else if x == "equal"{
			if ComesBefore(d[i].Person.LastName , great.Person.LastName) == "true"{
				d.swap(i, partition)
				partition++
			}
		}
	}

	d.swap(partition, high)
	return
}
func (d Datas) sortName(low int, high int) {
	if low < high {
		p := d.partitionName(low, high)
		d.sortName(low, p-1)
		d.sortName(p+1, high)
	}
}
func (d Datas) SortByName(){
	d.sortName(0, len(d)-1)
}
//Sort By Last Name
func (d Datas) partitionLastName(low int, high int) (partition int) {
	partition = low
	great := d[high]
	for i := low; i < high; i++ {

		if x := ComesBefore(d[i].Person.LastName , great.Person.LastName); x == "true" {
			d.swap(i, partition)
			partition++
		}else if x == "equal"{
			if ComesBefore(d[i].Person.Name , great.Person.Name) == "true"{
				d.swap(i, partition)
				partition++
			}
		}
	}

	d.swap(partition, high)
	return
}
func (d Datas) sortLastName(low int, high int) {
	if low < high {
		p := d.partitionLastName(low, high)
		d.sortLastName(low, p-1)
		d.sortLastName(p+1, high)
	}
}
func (d Datas) SortByLastName(){
	d.sortLastName(0, len(d)-1)
}