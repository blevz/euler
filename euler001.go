package main

import (
	"fmt"
	"github.com/blevz/euler/common"
)

func multiplesOfThreeAndFiveBelow(maxNum int, output chan<- int) {
	for x := 1; x < maxNum; x++ {
		if x%3 == 0 || x%5 == 0 {
			output <- x
		}
	}
	close(output)
}

func main() {
	c := make(chan int)
	go multiplesOfThreeAndFiveBelow(1000, c)
	sum := common.Accumulate_int(c)
	fmt.Println("sum is ", sum)
}
