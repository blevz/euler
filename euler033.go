package main

import (
	"fmt"
	"math/big"
)

func CheckEquals(num, den, nNum, nDen int64) bool {
	if den == 0 || nDen == 0 {
		return false
	}
	a := big.NewRat(num, den)
	b := big.NewRat(nNum, nDen)
	//fmt.Println(a, b)
	if a.Cmp(b) == 0 {
		return true
	}
	return false
}

func gcd(a, b int64) int64 {
	for a != 0 {
		c := a
		a = b % a
		b = c
	}
	return b
}

func main() {

	final_num := int64(1)
	final_denom := int64(1)
	//Denom must be 2 dig
	for denom := int64(11); denom < 100; denom++ {
		//num must be 2 dig and less than denom (to keep between 0 and 1)
		for numer := int64(10); numer < denom; numer++ {
			var newNumer, newDenom int64
			if numer%10 == denom%10 && numer%10 != 0 {
				newNumer = numer / 10
				newDenom = denom / 10
				if CheckEquals(numer, denom, newNumer, newDenom) {
					fmt.Println(numer, denom)
					final_num *= numer
					final_denom *= denom
				}
			}
			if numer%10 == denom/10 {
				newNumer = numer / 10
				newDenom = denom % 10
				if CheckEquals(numer, denom, newNumer, newDenom) {
					fmt.Println(numer, denom)
					final_num *= numer
					final_denom *= denom
				}
			}
			if numer/10 == denom/10 {
				newNumer = numer % 10
				newDenom = denom % 10
				if CheckEquals(numer, denom, newNumer, newDenom) {
					fmt.Println(numer, denom)
					final_num *= numer
					final_denom *= denom
				}
			}
			if numer/10 == denom%10 {
				newNumer = numer % 10
				newDenom = denom / 10
				if CheckEquals(numer, denom, newNumer, newDenom) {
					fmt.Println(numer, denom)
					final_num *= numer
					final_denom *= denom
				}
			}
		}
	}
	g := gcd(final_num, final_denom)
	fmt.Println(final_num/g, final_denom/g)
}
