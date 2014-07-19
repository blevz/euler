package main

import (
	"fmt"
	"github.com/blevz/euler/common"
	_ "math/big"
	"strconv"
)

func accumulateFromStrings(input <-chan string) int64 {
	sum := int64(0)
	for x := range input {
		y, _ := strconv.Atoi(x)
		sum += int64(y)
	}
	return sum
}

func main() {
	c := make(chan string)
	d := make(chan string)
	path := "euler013_data.txt"
	go common.ReadLineByLine(path, c)

}
