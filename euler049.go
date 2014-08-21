package main

import (
	"fmt"
	"github.com/blevz/euler/common"
)

func IsPermutation(num1, num2 int) bool {
	dig := make([]int, 10, 10)
	for num1 != 0 {
		dig[num1%10]++
		num1 /= 10
	}
	for num2 != 0 {
		dig[num2%10]--
		num2 /= 10
	}
	for _, v := range dig {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	c := make(chan int)
	fourDigPrimes := make([]int, 0)
	go common.GenPrimesToNum(10000, c)
	for x := range c {
		if x > 1000 {
			fourDigPrimes = append(fourDigPrimes, x)
		}
	}
	for i, v1 := range fourDigPrimes {
		for _, v2 := range fourDigPrimes[i+1:] {
			if IsPermutation(v1, v2) {
				diff := v2 - v1
				v3 := v2 + diff
				if common.IsPrime(v3) && IsPermutation(v2, v3) {
					fmt.Println(v1, v2, v3)
				}
			}
		}
	}
}
