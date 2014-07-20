package main

import (
	"fmt"
)

type seenBefore struct {
	number    int
	remainder int
}

func getRecurrence(numerator int, denominator int) int {
	cur := numerator
	pastMap := make(map[seenBefore]int)
	for dig := 0; true; dig++ {
		cur *= 10
		num := cur / denominator
		rem := cur % denominator
		a := seenBefore{number: num, remainder: rem}
		val, exists := pastMap[a]
		if exists {
			return dig - val
		} else {
			pastMap[a] = dig
		}
		if rem == 0 {
			return 1
		}
		cur = rem
	}
	return 1
}

func main() {

	maxRecurrence := 0
	maxIndex := 0
	for x := 2; x < 1000; x++ {
		fmt.Println("1/", x)
		curRec := getRecurrence(1, x)
		if curRec > maxRecurrence {
			maxIndex = x
			maxRecurrence = curRec
		}
	}
	fmt.Println(maxIndex, " with ", maxRecurrence)
}
