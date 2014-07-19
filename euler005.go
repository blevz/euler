package main

import (
	"fmt"
	"github.com/blevz/euler/common"
)

func main() {
	product := 1
	for y := 1; y < 21; y++ {
		if common.IsPrime(y) {
			product *= y
		}
	}
	fmt.Println(product)
	product *= 8
	product *= 3
	for y := 1; y < 21; y++ {
		fmt.Println(y, " ", product%y)
	}
	fmt.Println(product)
}
