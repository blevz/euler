package main

import (
	"fmt"
	"github.com/blevz/euler/common"
)

func GenTriangleNums(numEntries int, output chan<- int) {
	val := 1
	for x := 0; x < numEntries; x++ {
		output <- val
		val += x + 2
	}
	close(output)
}

func main() {
	c := make(chan int)
	m := make(map[int]struct{})
	go GenTriangleNums(100, c)
	for x := range c {
		m[x] = struct{}{}
	}
	strArr, err := common.GetAllCSVStrings("euler042_data.txt")
	if err != nil {
		fmt.Println(err)
	}
	num := 0
	for _, x := range strArr {
		val := common.GetStringVal(x)
		_, exists := m[val]
		if exists {
			fmt.Println(x)
			num++
		}
	}
	fmt.Println(num)
}
