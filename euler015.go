package main

import (
	"fmt"
	"github.com/blevz/euler/common"
)

func main() {
	array := common.MakeTDArray(21, 21)
	for x := 0; x < 21; x++ {
		for y := 0; y < 21; y++ {
			numPaths := 1
			if x == 0 && y == 0 {
				numPaths = 1
			} else if x == 0 {
				numPaths = array.GetElemVal(x, y-1).(int)
			} else if y == 0 {
				numPaths = array.GetElemVal(x-1, y).(int)
			} else {
				numPaths = array.GetElemVal(x-1, y).(int) + array.GetElemVal(x, y-1).(int)
			}
			array.SetElem(x, y, numPaths)
		}
	}
	fmt.Println(array.GetElemVal(20, 20))
}
