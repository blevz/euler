package main

import (
	"fmt"
	"math/big"
)

func main() {
	m := make(map[string]struct{})
	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			bBig := big.NewInt(int64(a))
			result := big.NewInt(int64(a))
			for x := 1; x < b; x++ {
				result.Mul(result, bBig)
			}
			//fmt.Println(a, b, result)
			m[result.String()] = struct{}{}
		}
	}
	fmt.Println(len(m))
}
