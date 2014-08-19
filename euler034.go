package main

import (
	"fmt"
	"github.com/blevz/euler/common"
	_ "math"
)

func GetFactSum(num int) int {
	sum := 0
	for num != 0 {
		sum += common.Factorial(num % 10)
		num /= 10
	}
	return sum
}

func main() {
	sum := 0
	top := common.Factorial(9) * 10
	fmt.Println(top)
	for a := 3; a < top; a++ {
		if GetFactSum(a) == a {
			sum += a
			fmt.Println(a)
		}
	}
	fmt.Println(sum)
}
