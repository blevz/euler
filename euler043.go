package main

import (
	"fmt"
	"github.com/blevz/euler/common"
	"strconv"
)

func main() {
	primes := []int{0, 2, 3, 5, 7, 11, 13, 17}
	final := make([]int, 0)
	c := make(chan string)
	go common.ProducePermutations("0123456789", c)
	for x := range c {
		isDivisible := true
		for y := 1; y < len(x)-2; y++ {
			substr := x[y : y+3]
			num, _ := strconv.Atoi(substr)
			if num%primes[y] != 0 {
				isDivisible = false
			}
		}
		if isDivisible {
			num, _ := strconv.Atoi(x)
			final = append(final, num)
		}
	}
	sum := 0
	for _, x := range final {
		sum += x
	}
	fmt.Println(sum)
}
