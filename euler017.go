package main

import (
	"fmt"
	"strings"
)

var singleDigit = map[int]string{
	0: "",
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}
var teens = map[int]string{
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}

var tensDigit = map[int]string{
	0: "",
	1: "ten",
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

func int_string(num int) string {
	if num < 10 {
		return singleDigit[num]
	} else if num < 100 {
		tens := num / 10
		if tens == 1 {
			return teens[num]
		} else {
			return tensDigit[tens] + " " + singleDigit[num%10]
		}
	} else if num < 1000 {
		hundreds := num / 100
		if num%100 == 0 {
			return singleDigit[hundreds] + " hundred"
		} else {
			return singleDigit[hundreds] + " hundred and " + int_string(num%100)
		}
	} else {
		return "one thousand"
	}
	return ""
}

func main() {
	sum := 0
	for x := 1; x <= 1000; x++ {
		str := int_string(x)
		str = strings.Replace(str, " ", "", -1)
		sum += len(str)
		fmt.Println(str, len(str))
	}
	fmt.Println(sum)
}
