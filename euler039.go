package main

import (
	"fmt"
)

func main() {
	maxNum := 0
	maxP := 0
	for p := 3; p <= 1000; p++ {
		num := 0
		for c := 1; c <= p-2; c++ {
			for b := (p - c - 1); b <= p-1; b++ {
				a := p - c - b
				if a*a+b*b == c*c {
					num++
				}
			}
		}
		if num > maxNum {
			maxNum = num
			maxP = p
			fmt.Println(maxP, maxNum)
		}

	}

}
