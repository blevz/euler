package main

import (
	"fmt"
	"github.com/blevz/euler/common"
	"strconv"
)

func main() {

	maxNum := 10000000
	sl := make([]int, 0)

	for x := 2; x < maxNum; x++ {
		strNum := strconv.Itoa(x)
		sum := 0
		for _, c := range strNum {
			n := int(c - '0')
			sum += common.Pow(n, 5)
		}
		if sum == x {
			sl = append(sl, sum)
		}
	}
	fmt.Println(sl)
	sum := 0
	for _, x := range sl {
		sum += x
	}
	fmt.Println("sum is ", sum)
}
