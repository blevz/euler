package main

import (
	"fmt"
	"github.com/blevz/euler/common"
	"strconv"
)

func main() {
	c := make(chan string)
	go common.ProducePermutations("1234567", c)
	for x := range c {
		num, _ := strconv.Atoi(x)
		if common.IsPrime(num) {
			fmt.Println(num)
		}
	}
}
