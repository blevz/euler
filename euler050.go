package main

import (
	"fmt"
	"github.com/blevz/euler/common"
)

func main() {
	x := make(chan int)

	go common.GenPrimesToNum(1000000, x)
	num := 0
	maxNumBeforeMil := 0
	for v := range x {
		num += v
		maxNumBeforeMil++
		if num > 1000000 {
			break
		}
	}
	fmt.Println(maxNumBeforeMil)
	maxy := 21
	for y := maxy; y < maxNumBeforeMil; y++ {
		sum := 0
		lead := make(chan int)
		follow := make(chan int)
		go common.GenPrimesToNum(1000000, lead)
		go common.GenPrimesToNum(1000000, follow)

		for x := 0; x < y; x++ {
			sum += <-lead
		}
		for sum < 1000000 {
			if common.IsPrime(sum) {
				maxy = y
				fmt.Println(maxy, sum)
			}
			sum += <-lead
			sum -= <-follow
		}

		for b := range lead {
			b = b
		}
		for a := range follow {
			a = a
		}
		fmt.Println("run thru")
	}
}
