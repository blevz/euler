package main

import (
	"fmt"
)

func GetSpiralSum(size int) int {
	size = (size - 1) / 2
	sum := 1
	num := 1
	inc := 2
	for x := 0; x < size; x++ {
		for i := 0; i < 4; i++ {
			num += inc
			sum += num
		}
		inc += 2
	}
	return sum
}

func main() {
	spiral := GetSpiralSum(1001)
	fmt.Println(spiral)
}
