package main

import (
	"fmt"
	"github.com/blevz/euler/common"
)

//Returns true if none of the digits have been before
//Returns false if a zero is seen or a digit has been seen before
func checkDigits(num int, digits *[]bool) bool {
	for num != 0 {
		if num%10 == 0 {
			return false
		}
		v := (*digits)[num%10-1]
		if v {
			return false
		}
		(*digits)[num%10-1] = true
		num /= 10
	}
	return true
}

func AllTrue(digits []bool) bool {
	for _, v := range digits {
		if !v {
			return false
		}
	}
	return true
}

func main() {
	products := make(map[int]struct{})
	for a := 1; a < common.Pow(10, 8); a++ {
		aDigits := common.NumDigits(a)
		for b := a; b < common.Pow(10, 9-aDigits); b++ {
			bDigits := common.NumDigits(b)
			mul := a * b
			//If there are two many digits now, stop the inner loop
			//B cannot get any bigger and satisfy the constraints
			if bDigits+aDigits+common.NumDigits(mul) > 9 {
				break
			}

			digits := make([]bool, 9, 9)
			if (checkDigits(a, &digits) && checkDigits(b, &digits) && checkDigits(mul, &digits)) && AllTrue(digits) {
				products[mul] = struct{}{}
			}

		}
	}
	sum := 0
	for k, _ := range products {
		sum += k
	}
	fmt.Println(sum)
}
