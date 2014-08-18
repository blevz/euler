package main

import (
	"fmt"
	"github.com/blevz/euler/common"
	"os"
	"sort"
)

func GetStringVal(str string) int {
	sum := 0
	for _, r := range str {
		sum += int(r - 'A' + 1)
	}
	return sum
}

func main() {

	strArr := sort.StringSlice{}
	strArr, err := common.GetAllCSVStrings("euler022_data.txt")
	if err != nil {
		os.Exit(1)
	}
	strArr.Sort()
	sum := 0
	for i, v := range strArr {
		sum += (i + 1) * GetStringVal(v)
	}
	fmt.Println(sum)
}
