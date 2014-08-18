package main

import (
	"fmt"
	"github.com/blevz/euler/common"
)

func IsCircularPrime(num int) bool {
	numD := common.NumDigits(num)
	for x := 0; x < numD; x++ {
		num = (num / 10) + (num%10)*common.Pow(10, numD-1)
		//fmt.Println(num)
		if !common.IsPrime(num) {
			return false
		}
	}
	return true
}

func main() {
	num := 0
	c := make(chan int)
	go common.GenPrimesToNum(1000000, c)
	for x := range c {
		if IsCircularPrime(x) {
			fmt.Println(x)
			num++
		}
	}
	fmt.Println(num)
}
