package main

import (
	"fmt"
	"github.com/blevz/euler/common"
	"math/big"
)

func Is_Lychrel_Num(num *big.Int) bool {
	for i := 0; i <= 50; i++ {
		//fmt.Println(num, common.Reverse_BigNum(num))
		num.Add(num, common.Reverse_BigNum(num))
		if common.IsPalindrome(num.String()) {
			return true
		}
	}
	return false
}

func main() {
	num := 0
	Is_Lychrel_Num(big.NewInt(4994))
	for x := 1; x < 10000; x++ {
		if !Is_Lychrel_Num(big.NewInt(int64(x))) {
			num++
			//fmt.Println(x)
		}
	}
	fmt.Println(num)
}
