package main

import (
	"fmt"
	"github.com/blevz/euler/common"
)

func genFibToVal(maxNum int, output chan<- int) {
	a, b := 0, 1
	for b < maxNum {
		output <- b
		newB := a + b
		a = b
		b = newB
	}
	close(output)
}

func main() {
	c := make(chan int)
	d := make(chan int)
	go genFibToVal(4000000, c)
	go common.Remove_int(c, d, common.IsOdd)
	sum := common.Accumulate_int(d)
	fmt.Println("sum is ", sum)
}
