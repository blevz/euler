package main

import (
	"encoding/csv"
	"fmt"
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

	file, err := os.Open("euler022_data.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	strArr := sort.StringSlice{}
	reader := csv.NewReader(file)
	strArr, err = reader.Read()
	if err != nil {
		fmt.Println(err)
	}

	strArr.Sort()
	sum := 0
	for i, v := range strArr {
		sum += (i + 1) * GetStringVal(v)
	}
	fmt.Println(sum)
}
