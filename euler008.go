package main

import (
	"bufio"
	"fmt"
	"github.com/blevz/euler/common"
	"os"
)

func ReadLinesAndAppendStrings(path string) string {
	toReturn := ""
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		toReturn += scanner.Text()
	}
	return toReturn
}

func produceProductOfChar(str string, numChar int, output chan<- int) {

	for x := 0; x < len(str)-numChar; x++ {
		fmt.Println("test", str[x]-'0')
		product := 1
		for y := 0; y < numChar; y++ {
			product = product * int(str[y+x]-'0')
		}
		output <- product
	}
	close(output)
}

func main() {
	c := make(chan int)
	str := ReadLinesAndAppendStrings("euler008_data.txt")
	go produceProductOfChar(str, 13, c)
	max := common.Max_int(c)
	fmt.Println(max)
}
