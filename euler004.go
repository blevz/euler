package main

import (
	"fmt"
	"github.com/blevz/euler/common"
)

func fun(d chan<- int) {
	for x := 1; x < 1000; x++ {
		for y := 1; y < 1000; y++ {
			product := x * y
			if common.IsPalindromeNum(product) {
				d <- product
			}
		}
	}
	close(d)
}

func main() {
	c := make(chan int)
	go fun(c)

	fmt.Println(common.Max_int(c))
}
