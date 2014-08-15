package main

import (
	"fmt"
	"github.com/blevz/euler/common"
)

func numDigits(x int) int {
	numDig := 0
	for ; x != 0; x /= 10 {
		numDig++
	}
	return numDig
}

func testTrunct(num int, m map[int]struct{}) bool {
	for x := num; x != 0; x /= 10 {
		_, exists := m[x]
		if !exists {
			return false
		}
	}

	for x := num; x >= 1; x = x % common.Pow(10, numDigits(x)-1) {
		_, exists := m[x]
		if !exists {
			return false
		}
	}

	return true
}

func main() {
	m := make(map[int]struct{})
	trunctables := make([]int, 0)
	primes := make(chan int)
	go common.GenPrimes(primes)
	for x := range primes {
		m[x] = struct{}{}
		if testTrunct(x, m) && x > 10 {
			trunctables = append(trunctables, x)

		}
		if len(trunctables) == 11 {
			break
		}
	}
	fmt.Println(trunctables)
	sum := 0
	for _, x := range trunctables {
		sum += x
	}
	fmt.Println(sum)
}
