package common

import (
	"math/big"
)

func Reverse_Num(num int) int {
	toReturn := 0
	for num != 0 {
		toReturn = toReturn*10 + num%10
		num /= 10
	}
	return toReturn
}

func Reverse_BigNum(input *big.Int) *big.Int {
	num := big.NewInt(0)
	num.Set(input)
	toReturn := big.NewInt(0)
	ten := big.NewInt(10)
	for num.String() != "0" {
		temp := big.NewInt(0)
		temp.Mod(num, ten)
		toReturn.Mul(ten, toReturn)
		toReturn.Add(toReturn, temp)
		num.Div(num, ten)
	}
	return toReturn
}
