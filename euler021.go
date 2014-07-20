package main

import (
	"fmt"
	"github.com/blevz/euler/common"
)

func getDivisorList(num int) []int {
	toReturn := make([]int, 0)
	toReturn = append(toReturn, 1)
	for x := 2; x <= common.Sqrt(num); x++ {
		if num%x == 0 {
			toReturn = append(toReturn, x)
			toReturn = append(toReturn, num/x)
		}
	}
	return toReturn
}

func main() {
	dMap := make(map[int]int)
	for x := 1; x < 10000; x++ {
		sum := 0
		for _, v := range getDivisorList(x) {
			sum += v
		}
		dMap[x] = sum
	}
	sum := 0
	for x := 1; x < 10000; x++ {
		v, _ := dMap[x]
		z, _ := dMap[v]
		if x == z && x != v {
			fmt.Println(x, v)
			sum += x
		}
	}
	fmt.Println(sum)

}
