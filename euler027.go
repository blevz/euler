package main

import (
	"fmt"
	"github.com/blevz/euler/common"
)

func main() {
	max_n := 0
	for a := -999; a < 1000; a++ {
		for b := -999; b < 1000; b++ {
			for n := 0; true; n++ {
				num := n*n + a*n + b
				if num < 0 {
					break
				}
				if !common.IsPrimeWithDataStore(num) {
					if n > max_n {
						fmt.Println(a, b, n)
						fmt.Println(a * b)
						max_n = n
					}
					break
				}
			}
		}
	}

}
