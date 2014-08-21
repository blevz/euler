package main

import (
	"fmt"
	"github.com/blevz/euler/common"
	"strconv"
)

func Is19Pandigital(numstr string) bool {
	if len(numstr) != 9 {
		return false
	}
	digits := make([]bool, 9, 9)
	for _, r := range numstr {
		if r-'0' == 0 {
			return false
		}
		v := digits[int(r-'1')]
		if v {
			return false
		}
		digits[int(r-'1')] = true
	}
	return true
}

func CheckForPandigMulti(x int) string {
	numDigits := common.NumDigits(x)
	numStr := strconv.Itoa(x)
	for a := 2; numDigits < 9; a++ {
		nextNum := x * a
		numDigits += common.NumDigits(nextNum)
		numStr = numStr + strconv.Itoa(nextNum)
	}
	if Is19Pandigital(numStr) {
		fmt.Println(numStr)
		return numStr
	}
	return ""
}

func main() {
	for x := 90; x < 100; x++ {
		CheckForPandigMulti(x)
	}
	for x := 900; x < 1000; x++ {
		CheckForPandigMulti(x)
	}
	for x := 9000; x < 10000; x++ {
		CheckForPandigMulti(x)
	}
}
