package main

import (
	"fmt"
	"github.com/blevz/euler/common"
	"math"
)

func genNum(numMax int, output chan<- int) {
	for x := 2; x < numMax; x++ {
		output <- x
	}
	close(output)
}

func main() {
	xth := 10001
	c := make(chan int)
	go genNum(int(math.MaxInt32), c)
	y := 0
	for x := range c {
		if common.IsPrime(x) {
			y++
			if y == xth {
				fmt.Println(x)
				break
			}

		}
	}
}
