package main

import (
	"fmt"
	"github.com/blevz/euler/common"
	"math"
)

func Check(primes []int, dubSquare []int, x int) (bool, int, int) {
	for _, v1 := range primes {
		for _, v2 := range dubSquare {
			if v1+v2 == x {
				return true, v1, v2
			}
		}
	}
	return false, -1, -1
}

func main() {
	dubSquareBase := 1
	dubSquare := make([]int, 0)
	dubSquare = append(dubSquare, dubSquareBase*dubSquareBase*2)
	dubSquareBase++

	primes := make([]int, 0)
	primes = append(primes, 2, 3, 5, 7)
	for x := 9; x < int(math.MaxInt32); x += 2 {
		for x > dubSquare[len(dubSquare)-1] {
			dubSquare = append(dubSquare, dubSquareBase*dubSquareBase*2)
			dubSquareBase++
		}

		isPrime := common.IsPrime(x)
		if isPrime {
			primes = append(primes, x)
		} else {
			c, a, b := Check(primes, dubSquare, x)
			if !c {
				fmt.Println(x, a, b)
				return
			}
		}
	}
}
