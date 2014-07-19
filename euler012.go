package main

import (
	"fmt"
	"github.com/blevz/euler/common"
)

func genTriangleNums(output chan<- int) {
	tri := 0
	for x := 1; true; x++ {
		tri += x
		output <- tri
	}
}

func hasXDivisors(num int) int {
	mp := common.PrimeFactorizationMap(num)
	product := 1
	for _, freq := range mp {
		product *= freq + 1
	}
	return product
}

func main() {
	c := make(chan int)
	go genTriangleNums(c)
	for x := range c {
		if hasXDivisors(x) > 500 {
			fmt.Println(x)
			break
		}
	}
}
