package main

import "strings"

func ComesBefore(first string , second string) string{
	first = strings.ToLower(first)
	second = strings.ToLower(second)
	firstSize := len(first)
	secondSize := len(second)
	size := 0
	if firstSize < secondSize{
		size = firstSize
	}else{
		size = secondSize
	}
	if first == second{
		return "equal"
	}
	if strings.HasPrefix(first , second){
		return "false"
	}else if strings.HasPrefix(second , first){
		return "true"
	}else{
		for i := 0 ; i < size ; i++{
			if first[i] < second[i]{
				return "true"
			}else if first[i] > second[i]{
				return "false"
			}
		}
		return "equal"
	}

}
