package main

import (
	"fmt"
	"math"
)

func GetDigList(a int) []int {
	toReturn := make([]int, 10, 10)
	for a != 0 {
		toReturn[a%10]++
		a /= 10
	}
	return toReturn
}

func HasSameDigits(a, b int) bool {
	aDig := GetDigList(a)
	bDig := GetDigList(b)
	for i, v := range aDig {
		if bDig[i] != v {
			return false
		}
	}
	return true
}

func CheckPermutedMultiples(x int) bool {
	for mul := 2; mul <= 6; mul++ {
		if !HasSameDigits(x, x*mul) {
			return false
		}
	}
	return true
}

func main() {

	for x := 10; x < math.MaxInt32; x++ {
		if CheckPermutedMultiples(x) {
			fmt.Print(x)
			break
		}
	}

}
