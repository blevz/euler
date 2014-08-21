package main

import (
	"fmt"
	"github.com/blevz/euler/common"
)

func GetDigit(num int) int {
	for x := 1; ; x++ {
		num -= common.NumDigits(x)
		if num <= 0 {
			toReturn := x
			for ; num < 0; num++ {
				toReturn /= 10
			}
			return toReturn % 10
		}
	}
}

func main() {
	product := 1
	for x := 0; x < 7; x++ {
		y := common.Pow(10, x)
		product *= GetDigit(y)
		//fmt.Println(common.Pow(10, x), GetDigit(common.Pow(10, x)))
	}
	fmt.Println(product)
}
