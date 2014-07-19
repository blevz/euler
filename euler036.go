package main

import (
	"fmt"
	"github.com/blevz/euler/common"
	"strconv"
)

func main() {
	sum := 0
	for x := 0; x < 1000000; x++ {
		if common.IsPalindromeNum(x) {
			boolRep := strconv.FormatInt(int64(x), 2)
			if common.IsPalindrome(boolRep) {
				sum += x
				fmt.Println(x)
			}
		}
	}
	fmt.Println(sum)
}
