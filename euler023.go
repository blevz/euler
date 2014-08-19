package main

import (
	"fmt"
	"github.com/blevz/euler/common"
)

func IsAbundandant(k int) bool {
	sum := 1
	for x := 2; x < common.Sqrt(k)+1; x++ {
		if k%x == 0 {
			sum += x
			if k/x != x {
				sum += k / x
			}

		}
	}
	fmt.Println(k, sum)
	return k < sum
}

func SumOfTwoAbundFill(m map[int]struct{}, abund map[int]struct{}) int {
	for v1, _ := range m {
		for v2, _ := range m {
			abund[v1+v2] = struct{}{}
		}
	}
	sum := 0
	for v, _ := range abund {
		if v <= 28123 {
			sum += v
		}
	}
	return sum
}

func main() {
	y := make(map[int]struct{})
	m := make(map[int]struct{})
	total := 0
	for x := 1; x <= 28123; x++ {
		if IsAbundandant(x) {
			m[x] = struct{}{}
			fmt.Println(x)
		}
		total += x
	}
	fmt.Println(total)
	sum := SumOfTwoAbundFill(m, y)
	fmt.Println(sum)
	fmt.Println(total - sum)

}
