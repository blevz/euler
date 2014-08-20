package main

import (
	"fmt"
)

func GetNumWays(num int, pos []int) int {
	if num == 0 {
		return 1
	}
	sum := 0
	for i, v := range pos {
		for times := 1; num-v*times >= 0; times++ {
			sum += GetNumWays(num-v*times, pos[i+1:])
		}
	}
	return sum
}

func main() {
	//var possible []int = []int{1, 2, 5, 10, 20, 50, 100, 200}
	var possibleRev []int = []int{200, 100, 50, 20, 10, 5, 2, 1}
	fmt.Println(GetNumWays(200, possibleRev))
}
