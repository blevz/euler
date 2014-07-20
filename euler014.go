package main

import (
	"fmt"
)

var collatz = make(map[int]int)

func getCollatz(x int) int {
	num := 1
	for x != 1 {

		if x%2 == 0 {
			x = x / 2
		} else {
			x = (3 * x) + 1
		}

		c, exists := collatz[x]
		if exists {
			return num + c
		}

		num++
	}
	return num
}

func main() {
	collatz[1] = 1
	largestPath, largestIndex := 0, 0
	for x := 1; x < 1000000; x++ {
		v := getCollatz(x)
		if largestPath < v {
			largestPath = v
			largestIndex = x
		}
		collatz[x] = v
	}
	fmt.Println(largestIndex, ", ", largestPath)
}
