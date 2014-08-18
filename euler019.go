package main

import (
	"fmt"
)

type Month struct {
	Name    string
	numDays int
}

func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

var months = []Month{
	{"Januaryu", 31},
	{"February", 28},
	{"March", 31},
	{"April", 30},
	{"May", 31},
	{"June", 30},
	{"July", 31},
	{"August", 31},
	{"September", 30},
	{"October", 31},
	{"November", 30},
	{"December", 31},
}

var dayOfWeek = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednsday",
	"Thursday",
	"Friday",
	"Saturday",
}

func main() {
	curDay := 1
	numDaysInRegularYear := 0
	for _, month := range months {
		numDaysInRegularYear += month.numDays
	}
	count := 0
	//Make up for the first year
	curDay = (curDay + numDaysInRegularYear) % 7

	//Go thru each year
	for x := 1901; x <= 2000; x++ {
		for _, month := range months {
			curDay += month.numDays
			if month.Name == "February" && IsLeapYear(x) {
				curDay++
			}
			curDay = curDay % 7
			if curDay == 0 {
				//fmt.Println(month.Name, x)
				count++
			}
		}
	}
	fmt.Println(count)
}
