package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	file, err := os.Open("euler099_data.txt")
	if err != nil {
		fmt.Println(err)
	}
	var base, exp int
	lineNum := 0
	max := float64(0)
	for {
		lineNum++
		_, err := fmt.Fscanf(file, "%d,%d", &base, &exp)
		if err != nil {
			break
		}
		val := math.Log10(float64(base)) * float64(exp)
		if max <= val {
			max = val
			fmt.Println(lineNum, base, exp, val)
		}
	}

}
