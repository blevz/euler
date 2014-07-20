package main

import (
	"fmt"
	"github.com/blevz/euler/common"
	"math/big"
	_ "strconv"
)

func accumulateFromStrings(input <-chan string) *big.Int {
	sum := big.NewInt(0)
	for x := range input {
		z := big.NewInt(0)
		z.UnmarshalText([]byte(x))
		fmt.Println(z)
		sum.Add(sum, z)
	}
	return sum
}

func main() {
	c := make(chan string)
	path := "euler013_data.txt"
	go common.ReadLineByLine(path, c)
	sum := accumulateFromStrings(c)
	fmt.Println(sum)
	sumStr, _ := sum.MarshalText()
	fmt.Println(string(sumStr[0:10]))
}
