package main

import (
	"encoding/csv"
	_ "errors"
	"fmt"
	"github.com/blevz/euler/common"
	"os"
	"strconv"
)

func CSVToTwoDArray(filename string, xSize, ySize int) (*common.TDArray, error) {
	toReturn := common.MakeTDArray(xSize, ySize)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	reader := csv.NewReader(file)
	for y := 0; y < ySize; y++ {
		a, _ := reader.Read()
		for x := 0; x < xSize; x++ {
			n, _ := strconv.Atoi(a[x])
			toReturn.SetElem(x, y, n)
		}
	}
	return toReturn, nil
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	v, err := CSVToTwoDArray("euler081_data.txt", 80, 80)
	if err != nil {
		fmt.Println(err)
	}
	for x := 79; x >= 0; x-- {
		for y := 79; y >= 0; y-- {
			var minPath int
			if x == 79 && y == 79 {
				minPath = v.GetElemVal(x, y).(int)
			} else if x == 79 {
				minPath = v.GetElemVal(x, y).(int) + v.GetElemVal(x, y+1).(int)
			} else if y == 79 {
				minPath = v.GetElemVal(x, y).(int) + v.GetElemVal(x+1, y).(int)
			} else {
				minPath = v.GetElemVal(x, y).(int) + min(v.GetElemVal(x+1, y).(int), v.GetElemVal(x, y+1).(int))
			}
			v.SetElem(x, y, minPath)
		}
	}
	fmt.Println(v.GetElemVal(0, 0))
}
