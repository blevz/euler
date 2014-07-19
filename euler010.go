package main

import (
	"fmt"
	"github.com/blevz/euler/common"
)

func genAllPrimesLessThanNum(maxNum int, output chan<- int) {
	for x := 2; x < maxNum; x++ {
		if common.IsPrime(x) {
			output <- x
		}
	}
	close(output)
}

func main() {
	c := make(chan int)
	go genAllPrimesLessThanNum(2000000, c)
	sum := common.Accumulate_int(c)
	fmt.Println(sum)
}
